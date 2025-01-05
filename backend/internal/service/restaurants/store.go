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
	UpdateRestaurantProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error
	UpdateRestaurantIsActive(ctx context.Context, userID uuid.UUID, isActive bool) error
	GetRestaurantPublicProfile(ctx context.Context, id uuid.UUID) (*models.Restaurant, error)
	GetAllRestaurantPublicProfile(ctx context.Context) (*models.RestaurantPublicProfileList, error)
	GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error)
}

type RestaurantRepo struct {
	client *gorm.DB
}

func NewRestaurantRepo(client *gorm.DB) *RestaurantRepo {
	return &RestaurantRepo{client: client}
}

func (r *RestaurantRepo) UpdateRestaurant(ctx context.Context, userID uuid.UUID, restaurant *models.RestaurantUpdate) error {
	return r.client.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).Updates(restaurant).Error
}

func (r *RestaurantRepo) UpdateRestaurantProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error {
	return r.client.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).Update("profile_img", filePath).Error
}

func (r *RestaurantRepo) UpdateRestaurantIsActive(ctx context.Context, userID uuid.UUID, isActive bool) error {
	return r.client.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).Update("is_active", isActive).Error
}

func (r *RestaurantRepo) GetRestaurantPublicProfile(ctx context.Context, id uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := r.client.WithContext(ctx).Model(&models.Restaurant{}).Where("id = ?", id).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (r *RestaurantRepo) GetAllRestaurantPublicProfile(ctx context.Context) (*models.RestaurantPublicProfileList, error) {
	var res []models.Restaurant
	var restaurants models.RestaurantPublicProfileList
	if err := r.client.WithContext(ctx).Model(&models.Restaurant{}).Find(&res).Scan(&restaurants).Error; err != nil {
		return nil, err
	}
	return &restaurants, nil
}

func  GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := config.Session.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (r *RestaurantRepo)  GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := config.Session.WithContext(ctx).Model(&models.Restaurant{}).Where("user_id = ?", userID).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
