package categories

import "gorm.io/gorm"

type CategoryStore interface {
}

type CategoryRepo struct {
	client *gorm.DB
}

func NewCategoryRepo(client *gorm.DB) *CategoryRepo {
	return &CategoryRepo{client: client}
}
