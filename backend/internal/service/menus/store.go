package menus

import (
	"context"

	"github.com/JerryJeager/JeagerEats/config"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"gorm.io/gorm"
)

type MenuStore interface {
	CreateMenu(ctx context.Context, menu *models.Menu) error
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
