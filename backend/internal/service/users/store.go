package users

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User, restaurant *models.Restaurant, rider *models.Rider) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)

	GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User, restaurant *models.Restaurant, rider *models.Rider) error {

	if user.Role == models.CUSTOMER {
		return r.client.Create(user).WithContext(ctx).Error
	} else {
		err := r.client.Transaction(func(tx *gorm.DB) error {
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
	if err := r.client.Where("email = ?", email).First(&user).WithContext(ctx).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.client.Where("id = ?", userID).First(&user).WithContext(ctx).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := r.client.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
