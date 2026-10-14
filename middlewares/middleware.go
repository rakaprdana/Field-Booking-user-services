package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	response "user-services/common/reponses"
	"user-services/config"
	"user-services/constants"
	constErr "user-services/constants/errors"
	services "user-services/services/users"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("Recoverd from panic %v", r)
				ctx.JSON(http.StatusInternalServerError, response.Response{
					Status:  constants.Error,
					Message: constErr.ErrInternalServerError.Error(),
				})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

func RateLimiter(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusTooManyRequests, response.Response{
				Status:  constants.Error,
				Message: constErr.ErrTooManyRequests.Error(),
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func extactBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")

	if len(arrayToken) == 2 {
		return arrayToken[1]
	}
	return ""
}

func responseUnauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, response.Response{
		Status:  constants.Error,
		Message: message,
	})
	ctx.Abort()
}

func validateAPIKey(ctx *gin.Context) error {
	apiKey := ctx.GetHeader(constants.XApiKey)
	requestAt := ctx.GetHeader(constants.XRequestAt)
	serivceName := ctx.GetHeader(constants.XServiceName)
	signatureKey := config.Config.SignatureKey

	validateKey := fmt.Sprintf("%s:%s:%s", serivceName, signatureKey, requestAt)

	hash := sha256.New()
	hash.Write([]byte(validateKey))

	resulHash := hex.EncodeToString(hash.Sum(nil))
	if apiKey != resulHash {
		return constErr.ErrUnauthorized
	}

	return nil
}

func validateBearerToken(ctx *gin.Context, token string) error {
	if !strings.Contains(token, "Bearer") {
		return constErr.ErrUnauthorized
	}

	tokenString := extactBearerToken(token)
	if tokenString == "" {
		return constErr.ErrUnauthorized
	}

	claims := &services.Claims{}
	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, constErr.ErrInvalidToken
		}

		jwtSecret := []byte(config.Config.JwtSecretKey)
		return jwtSecret, nil
	})

	if err != nil || !tokenJwt.Valid {
		return constErr.ErrUnauthorized
	}

	userLogin := ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), constants.UserLogin, claims.User))
	ctx.Request = userLogin
	ctx.Set(constants.Token, token)
	return nil
}

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		token := ctx.GetHeader(constants.Authorization)
		if token == "" {
			responseUnauthorized(ctx, constErr.ErrUnauthorized.Error())
			return
		}

		err = validateBearerToken(ctx, token)
		if err != nil {
			responseUnauthorized(ctx, err.Error())
			return
		}

		err = validateAPIKey(ctx)
		if err != nil {
			responseUnauthorized(ctx, err.Error())
			return
		}
		ctx.Next()
	}
}
