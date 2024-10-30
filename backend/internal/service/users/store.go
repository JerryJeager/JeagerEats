package users

import (
	"context"

	"github.com/JerryJeager/JeagerEats/config"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User, restaurant *models.Restaurant, rider *models.Rider) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User, restaurant *models.Restaurant, rider *models.Rider) error {

	if user.Role == models.CUSTOMER {
		return config.Session.Create(user).WithContext(ctx).Error
	} else {
		err := config.Session.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(user).WithContext(ctx).Error; err != nil {
				return err
			}
			switch user.Role {
			case models.VENDOR:
				return tx.Create(restaurant).WithContext(ctx).Error
			case models.RIDER:
				return tx.Create(rider).WithContext(ctx).Error
			}
			return nil
		})
		return err
	}
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := config.Session.Where("email = ?", email).First(&user).WithContext(ctx).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
