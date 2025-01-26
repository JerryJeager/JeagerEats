package orders

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"gorm.io/gorm"
)

type OrderStore interface {
	CreateOrder(ctx context.Context, order *models.Order, items *[]models.Item) error
}

type OrderRepo struct {
	client *gorm.DB
}

func NewOrderRepo(client *gorm.DB) *OrderRepo {
	return &OrderRepo{client: client}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, order *models.Order, items *[]models.Item) error {

	err := r.client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).WithContext(ctx).Error; err != nil {
			return err
		}
		for _, item := range *items {
			if err := tx.Create(&item).WithContext(ctx).Error; err != nil {
				return err
			}
			if err := tx.Model(&models.Menu{}).Where("id = ?", item.MenuID).Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
