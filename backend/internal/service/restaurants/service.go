package restaurants

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type RestaurantSv interface {
	UpdateRestaurant(ctx context.Context, userID uuid.UUID, restaurant *models.RestaurantUpdate) error
	UpdateRestaurantProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error
	UpdateRestaurantIsActive(ctx context.Context, userID uuid.UUID, isActive *models.IsActive) error
}

type RestaurantServ struct {
	repo RestaurantStore
}

func NewRestaurantService(repo RestaurantStore) *RestaurantServ {
	return &RestaurantServ{repo: repo}
}

func (s *RestaurantServ) UpdateRestaurant(ctx context.Context, userID uuid.UUID, restaurant *models.RestaurantUpdate) error {
	return s.repo.UpdateRestaurant(ctx, userID, restaurant)
}

func (s *RestaurantServ) UpdateRestaurantProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error {
	return s.repo.UpdateRestaurantProfileImg(ctx, userID, filePath)
}

func (s *RestaurantServ) UpdateRestaurantIsActive(ctx context.Context, userID uuid.UUID, isActive *models.IsActive) error {
	return s.repo.UpdateRestaurantIsActive(ctx, userID, isActive.IsActive)
}
