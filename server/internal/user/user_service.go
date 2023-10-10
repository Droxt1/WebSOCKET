package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"server/util"
	"strconv"
	"time"
)

const (
	secrestKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repo Repository) Service {
	return &service{
		repo,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, userRequest *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: hashedPassword,
	}

	createdUser, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		ID:       strconv.Itoa(int(createdUser.ID)),
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}

	return res, nil

}

type MyJWTClaims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.RegisteredClaims
}

func (s *service) LoginUser(c context.Context, userRequest *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, userRequest.Email)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	err = util.CheckPasswordHash(userRequest.Password, user.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	// signed string as an access token
	ss, err := token.SignedString([]byte(secrestKey))
	if err != nil {
		return &LoginUserResponse{}, err
	}
	return &LoginUserResponse{accessToken: ss, Username: user.Username, ID: strconv.Itoa(int(user.ID))}, nil
}
