package repositories

import (
	"context"

	"github.com/handarudwiki/mkp-tht/internal/models"
	"gorm.io/gorm"
)

type (
	Terminal interface {
		Create(ctx context.Context, user models.Terminal) (models.Terminal, error)
	}

	terminalRepository struct {
		db *gorm.DB
	}
)

func NewTerminalRepository(db *gorm.DB) Terminal {
	return &terminalRepository{
		db: db,
	}
}

func (r *terminalRepository) Create(ctx context.Context, terminal models.Terminal) (models.Terminal, error) {
	err := r.db.WithContext(ctx).Create(&terminal).Error
	if err != nil {
		return terminal, err
	}
	return terminal, nil
}
