package menus

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type MenuSv interface {
	CreateMenu(ctx context.Context, restaurantID uuid.UUID, menu *models.Menu) (string, error)
	UpdateMenuImage(ctx context.Context, menuID uuid.UUID, filePath string) error
	GetMenusByRestaurantID(ctx context.Context, restaurantID uuid.UUID) (*models.Menus, error)
	GetMenuByID(ctx context.Context, menuID uuid.UUID) (*models.Menu, error)
	GetMenus(ctx context.Context) (*models.Menus, error)
	DeleteMenu(ctx context.Context, menuID uuid.UUID) error
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

func (s *MenuServ) UpdateMenuImage(ctx context.Context, menuID uuid.UUID, filePath string) error {
	return s.repo.UpdateMenuImage(ctx, menuID, filePath)
}

func (s *MenuServ) GetMenusByRestaurantID(ctx context.Context, restaurantID uuid.UUID) (*models.Menus, error) {
	return s.repo.GetMenusByRestaurantID(ctx, restaurantID)
}

func (s *MenuServ) GetMenuByID(ctx context.Context, menuID uuid.UUID) (*models.Menu, error) {
	return s.repo.GetMenuByID(ctx, menuID)
}

func (s *MenuServ) GetMenus(ctx context.Context) (*models.Menus, error) {
	return s.repo.GetMenus(ctx)
}	

func (s *MenuServ) DeleteMenu(ctx context.Context, menuID uuid.UUID) error {
	return s.repo.DeleteMenu(ctx, menuID)
}

