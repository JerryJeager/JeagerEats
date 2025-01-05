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
	GetRestaurantPublicProfile(ctx context.Context, id uuid.UUID) (*models.RestaurantPublicProfile, error)
	GetAllRestaurantPublicProfile(ctx context.Context) (*models.RestaurantPublicProfileList, error)
	GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error)
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

func (s *RestaurantServ) GetRestaurantPublicProfile(ctx context.Context, id uuid.UUID) (*models.RestaurantPublicProfile, error) {
	r, err := s.repo.GetRestaurantPublicProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.RestaurantPublicProfile{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Address:     r.Address,
		ProfileImg:  r.ProfileImg,
		Rating:      r.Rating,
		IsActive:    r.IsActive,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		CuisineType: r.CuisineType,
	}, nil
}

func (s *RestaurantServ) GetAllRestaurantPublicProfile(ctx context.Context) (*models.RestaurantPublicProfileList, error) {
	return s.repo.GetAllRestaurantPublicProfile(ctx)
}

func (s *RestaurantServ) GetRestaurant(ctx context.Context, userID uuid.UUID) (*models.Restaurant, error){
	return s.repo.GetRestaurant(ctx, userID)
}