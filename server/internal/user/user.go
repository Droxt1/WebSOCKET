package user

import (
	"context"
	"errors"
	"server/util"
	"strings"
)

// TODO: use GORM
type User struct {
	ID       int64  `json:"id" db:"id" uuid:"true"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("username is required")
	}

	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is required")
	}

	validEmail, _ := util.IsEmailValid(u.Email)
	if !validEmail {
		return errors.New("invalid email format")

	}

	validUsername, _ := util.IsUsernameValid(u.Username)
	if !validUsername {
		return errors.New("invalid username format")
	}

	if strings.TrimSpace(u.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

// Repository interface is used to define the methods that we want to use in our service
// the context is used to pass the context to the repository so we can use it to cancel the request if needed
type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

// Service interface is used to define the methods that we want to use in our handler
// It's basically a mere implementation for the business logic of the application
// and it's used to separate the business logic from the transport layer
type Service interface {
	CreateUser(c context.Context, user *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(c context.Context, user *LoginUserRequest) (*LoginUserResponse, error)
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserResponse struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
type LoginUserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type LoginUserResponse struct {
	accessToken string
	Username    string `json:"username" db:"username"`
	ID          string `json:"id" db:"id"`
}
