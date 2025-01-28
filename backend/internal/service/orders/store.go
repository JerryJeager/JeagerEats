package orders

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStore interface {
	CreateOrder(ctx context.Context, order *models.Order, items *[]models.Item) error

	GetMenu(ctx context.Context, menuID uuid.UUID) (*models.Menu, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)

	GetRiders(ctx context.Context) (*[]models.User, error)
	GetRestaurant(ctx context.Context, restaurantID uuid.UUID) (*models.Restaurant, error)
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

func (r *OrderRepo) GetMenu(ctx context.Context, menuID uuid.UUID) (*models.Menu, error) {
	var menu models.Menu
	if err := r.client.WithContext(ctx).Where("id = ?", menuID).First(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *OrderRepo) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.client.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *OrderRepo) GetRiders(ctx context.Context) (*[]models.User, error) {
	var riders []models.User
	if err := r.client.WithContext(ctx).Where("role = ?", models.RIDER).Find(&riders).Error; err != nil {
		return nil, err
	}
	return &riders, nil
}

func (r *OrderRepo) GetRestaurant(ctx context.Context, restaurantID uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := r.client.WithContext(ctx).Where("id = ?", restaurantID).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
