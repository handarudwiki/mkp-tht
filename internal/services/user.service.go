package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/mkp-tht/internal/dto"
	"github.com/handarudwiki/mkp-tht/internal/models"
	commons "github.com/handarudwiki/mkp-tht/internal/models/common"
	"github.com/handarudwiki/mkp-tht/internal/repositories"
	"github.com/handarudwiki/mkp-tht/internal/utils"
)

type (
	User interface {
		Register(ctx context.Context, dto dto.Register) (user models.User, errr error)
		Login(ctx context.Context, dto dto.Register) (token string, err error)
	}
	userService struct {
		repository repositories.User
		jwtSecret  string
	}
)

func NewUserService(repository repositories.User, jwtSecret string) User {
	return &userService{
		repository: repository,
		jwtSecret:  jwtSecret,
	}
}

func (s *userService) Register(ctx context.Context, dto dto.Register) (user models.User, err error) {
	user, err = s.repository.FindByEmail(ctx, dto.Email)

	if err != nil && !errors.Is(err, commons.ErrNotfound) {
		return
	}
	if user.ID != 0 {
		return
	}
	user = models.User{
		Email:    dto.Email,
		Password: dto.Password,
		Name:     dto.Name,
	}

	err = user.EncryptPassword()
	if err != nil {
		return
	}
	user, err = s.repository.Create(ctx, user)
	if err != nil {
		return
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, dto dto.Register) (token string, err error) {
	user, err := s.repository.FindByEmail(ctx, dto.Email)
	if err != nil {
		return "", commons.ErrCredentials
	}

	isValid, err := user.ComparePassword(dto.Password)
	if err != nil || !isValid {
		return "", commons.ErrWrongPassword
	}

	token, err = utils.GenerateToken(user.ID, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
