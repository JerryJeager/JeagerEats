package restaurants

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type RestaurantSv interface {
	 UpdateRestaurant(ctx context.Context, userID uuid.UUID, restaurant *models.RestaurantUpdate) error
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
