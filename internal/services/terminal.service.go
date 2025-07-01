package services

import (
	"context"

	"github.com/handarudwiki/mkp-tht/internal/dto"
	"github.com/handarudwiki/mkp-tht/internal/models"
	"github.com/handarudwiki/mkp-tht/internal/repositories"
)

type (
	Terminal interface {
		Create(ctx context.Context, dto dto.CreateTerminal) (models.Terminal, error)
	}

	terminalService struct {
		repository repositories.Terminal
	}
)

func NewTerminalService(repository repositories.Terminal) Terminal {
	return &terminalService{
		repository: repository,
	}
}

func (s *terminalService) Create(ctx context.Context, dto dto.CreateTerminal) (models.Terminal, error) {
	terminal := models.Terminal{
		Name:      dto.Name,
		Location:  dto.Location,
		Address:   dto.Address,
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
	}

	return s.repository.Create(ctx, terminal)
}
