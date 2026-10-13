package repositories

import (
	"context"
	"errors"
	commonErr "user-services/common/errors"
	constErr "user-services/constants/errors"
	"user-services/domain/dto"
	"user-services/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// FindByEmail implements [IUserRepository].
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constErr.ErrUserNotFound
		}
		return nil, commonErr.WrapError(constErr.ErrSQLError)
	}

	return &user, nil
}

// FindByUUID implements [IUserRepository].
func (r *UserRepository) FindByUUID(ctx context.Context, uuid string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("uuid = ?", uuid).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constErr.ErrUserNotFound
		}
		return nil, commonErr.WrapError(constErr.ErrSQLError)
	}

	return &user, nil
}

// FindByUsername implements [IUserRepository].
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constErr.ErrUserNotFound
		}
		return nil, commonErr.WrapError(constErr.ErrSQLError)
	}

	return &user, nil
}

// Register implements [IUserRepository].
func (r *UserRepository) Register(ctx context.Context, req *dto.RegisterRequest) (*models.User, error) {
	user := models.User{
		UUID:        uuid.New(),
		Name:        req.Name,
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		RoleID:      req.RoleID,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, commonErr.WrapError(constErr.ErrSQLError)
	}

	return &user, nil
}

// Update implements [IUserRepository].
func (r *UserRepository) Update(ctx context.Context, req *dto.UpdateUserRequest, uuid string) (*models.User, error) {
	user := models.User{
		Name:        req.Name,
		Username:    req.Username,
		Email:       req.Email,
		Password:    *req.Password,
		PhoneNumber: req.PhoneNumber,
	}

	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).Updates(&user).Error

	if err != nil {
		return nil, commonErr.WrapError(constErr.ErrSQLError)
	}

	return &user, nil
}

type IUserRepository interface {
	Register(context.Context, *dto.RegisterRequest) (*models.User, error)
	Update(context.Context, *dto.UpdateUserRequest, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUUID(context.Context, string) (*models.User, error)
}

// Constructor
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
