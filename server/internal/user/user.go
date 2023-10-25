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
	// Check if username is empty
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("username is required")
	}

	// Check if email is empty
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is required")
	}

	// Check if email format is valid
	validEmail, _ := util.IsEmailValid(u.Email)
	if !validEmail {
		return errors.New("invalid email format")

	}

	// Check if password is empty
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
