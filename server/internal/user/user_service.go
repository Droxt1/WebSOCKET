package user

import (
	"context"
	"server/util"
	"strconv"
	"time"
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

	return &CreateUserResponse{
		ID:       strconv.Itoa(int(createdUser.ID)),
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}, nil

}
