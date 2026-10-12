package errors

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrSQLError            = errors.New("Database server failed to execute query")
	ErrTooManyRequests     = errors.New("Too many requests")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrInvalidToken        = errors.New("Invalid token")
	ErrForbidden           = errors.New("Forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
