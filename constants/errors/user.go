package errors

import "errors"

var (
	ErrUserNotFound        = errors.New("User not found")
	ErrPasswordIncorrect   = errors.New("Password is incorrect")
	ErrUsernameExists      = errors.New("Username already exists")
	ErrEmailExists         = errors.New("Email already exists")
	ErrPaswordDoesNotMatch = errors.New("Password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExists,
	ErrPaswordDoesNotMatch,
}
