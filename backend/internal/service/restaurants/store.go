package restaurants

import (
	"context"

	"github.com/JerryJeager/JeagerEats/config"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestaurantStore interface {
	UpdateRestaurant(ctx context.Context, userID uuid.UUID, restaurant *models.RestaurantUpdate) error
}

type RestaurantRepo struct {
	client *gorm.DB
}

func NewRestaurantRepo(client *gorm.DB) *RestaurantRepo {
	return &RestaurantRepo{client: client}
}

func (r *RestaurantRepo) UpdateRestaurant(ctx context.Context, userID uuid.UUID, restaurant *models.RestaurantUpdate) error {
	return config.Session.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).Updates(restaurant).Error
}
