package categories

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"gorm.io/gorm"
)

type CategoryStore interface {
	CreateCategory(ctx context.Context, category *models.Category) error
}

type CategoryRepo struct {
	client *gorm.DB
}

func NewCategoryRepo(client *gorm.DB) *CategoryRepo {
	return &CategoryRepo{client: client}
}

func (r *CategoryRepo) CreateCategory(ctx context.Context, category *models.Category) error {
	return r.client.WithContext(ctx).Create(category).Error
}
