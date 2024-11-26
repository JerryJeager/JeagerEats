package categories

import (
	"context"
	"strings"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type CategorySv interface {
	CreateCategory(ctx context.Context, category *models.Category) (string, error)
}

type CategoryServ struct {
	repo CategoryStore
}

func NewCategoryService(repo CategoryStore) *CategoryServ {
	return &CategoryServ{repo: repo}
}

func (s *CategoryServ) CreateCategory(ctx context.Context, category *models.Category) (string, error) {
	id := uuid.New()
	category.ID = id
	category.Name = strings.ToLower(category.Name)
	if err := s.repo.CreateCategory(ctx, category); err != nil {
		return "", err
	}
	return id.String(), nil
}
