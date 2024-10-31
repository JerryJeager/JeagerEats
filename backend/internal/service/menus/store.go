package menus

import (
	"context"

	"github.com/JerryJeager/JeagerEats/config"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuStore interface {
	CreateMenu(ctx context.Context, menu *models.Menu) error
	UpdateMenuImage(ctx context.Context, menuID uuid.UUID, filePath string) error
}

type MenuRepo struct {
	client *gorm.DB
}

func NewMenuRepo(client *gorm.DB) *MenuRepo {
	return &MenuRepo{client: client}
}

func (r *MenuRepo) CreateMenu(ctx context.Context, menu *models.Menu) error {
	return config.Session.WithContext(ctx).Create(menu).Error
}

func (r *MenuRepo) UpdateMenuImage(ctx context.Context, menuID uuid.UUID, filePath string) error {
	return config.Session.WithContext(ctx).Model(&models.Menu{}).Where("id = ?", menuID).Update("img_url", filePath).Error
}
