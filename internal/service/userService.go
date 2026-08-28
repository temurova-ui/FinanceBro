package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/temurova-ui/FinanceBro/internal/model"
	"github.com/temurova-ui/FinanceBro/pkg/jwt"
)

type UserRepo interface {
	Create(ctx context.Context, user model.User)(int, error);

}

type UserService struct{
	repo UserRepo
}

func NewUserService (repo UserRepo)*UserService{
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Register (ctx context.Context, req model.Register)(int, error){
	if len(req.Name) == 0{
		return 0, errors.New("Name should't be empty")
	}

	if len(req.Email) == 0{
		return 0, errors.New("Email should't be empty")
	}
	
	if len(req.PasswordHash) == 0{
		return 0, errors.New("PasswordHash should't be empty")
	}

	hash, err := jwt.HashPassword(req.PasswordHash)
	if err != nil{
		return 0, fmt.Errorf("jwt.HashPassword: %w", err)
	}
	return u.repo.Create(ctx, model.User{
		Name: req.Name,
		Email: req.Email,
		PasswordHash: hash,
	})

}