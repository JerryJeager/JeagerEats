package menus

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type MenuSv interface {
	CreateMenu(ctx context.Context, restaurantID uuid.UUID, menu *models.Menu) (string, error)
}

type MenuServ struct {
	repo MenuStore
}

func NewMenuService(repo MenuStore) *MenuServ {
	return &MenuServ{repo: repo}
}

func (s *MenuServ) CreateMenu(ctx context.Context, restaurantID uuid.UUID, menu *models.Menu) (string, error) {
	id := uuid.New()
	menu.ID = id
	menu.RestaurantID = restaurantID
	if err := s.repo.CreateMenu(ctx, menu); err != nil {
		return "", err
	}
	return id.String(), nil
}
