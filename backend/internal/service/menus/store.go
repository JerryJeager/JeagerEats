package menus

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuStore interface {
	CreateMenu(ctx context.Context, menu *models.Menu) error
	UpdateMenuImage(ctx context.Context, menuID uuid.UUID, filePath string) error
	GetMenusByRestaurantID(ctx context.Context, restaurantID uuid.UUID) (*models.Menus, error)
	GetMenuByID(ctx context.Context, menuID uuid.UUID) (*models.Menu, error)
	GetMenus(ctx context.Context) (*models.Menus, error)
	DeleteMenu(ctx context.Context, menuID uuid.UUID) error
	UpdateMenu(ctx context.Context, menuID uuid.UUID, menu *models.MenuUpdate) error
}

type MenuRepo struct {
	client *gorm.DB
}

func NewMenuRepo(client *gorm.DB) *MenuRepo {
	return &MenuRepo{client: client}
}

func (r *MenuRepo) CreateMenu(ctx context.Context, menu *models.Menu) error {
	return r.client.WithContext(ctx).Create(menu).Error
}

func (r *MenuRepo) UpdateMenuImage(ctx context.Context, menuID uuid.UUID, filePath string) error {
	return r.client.WithContext(ctx).Model(&models.Menu{}).Where("id = ?", menuID).Update("img_url", filePath).Error
}

func (r *MenuRepo) GetMenusByRestaurantID(ctx context.Context, restaurantID uuid.UUID) (*models.Menus, error) {
	var menus models.Menus
	if err := r.client.WithContext(ctx).Where("restaurant_id = ?", restaurantID).Find(&menus).Error; err != nil {
		return nil, err
	}
	return &menus, nil
}

func (r *MenuRepo) GetMenuByID(ctx context.Context, menuID uuid.UUID) (*models.Menu, error) {
	var menu models.Menu
	if err := r.client.WithContext(ctx).Where("id = ?", menuID).First(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *MenuRepo) GetMenus(ctx context.Context) (*models.Menus, error) {
	var menus models.Menus
	if err := r.client.WithContext(ctx).Find(&menus).Error; err != nil {
		return nil, err
	}
	return &menus, nil
}

func (r *MenuRepo) DeleteMenu(ctx context.Context, menuID uuid.UUID) error {
	return r.client.WithContext(ctx).Delete(&models.Menu{}, menuID).Error
}

func (r *MenuRepo) UpdateMenu(ctx context.Context, menuID uuid.UUID, menu *models.MenuUpdate) error {
	return r.client.WithContext(ctx).Model(&models.Menu{}).Where("id = ?", menuID).Updates(menu).Error
}
