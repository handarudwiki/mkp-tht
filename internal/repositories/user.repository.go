package repositories

import (
	"context"

	"github.com/handarudwiki/mkp-tht/internal/models"
	"gorm.io/gorm"
)

type (
	User interface {
		FindByEmail(ctx context.Context, email string) (models.User, error)
		Create(ctx context.Context, user models.User) (models.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) User {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}

	user.Password = "" // Clear password before returning
	return user, nil
}
