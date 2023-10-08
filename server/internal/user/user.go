package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
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
