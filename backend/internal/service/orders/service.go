package orders

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type OrderSv interface {
	CreateOrder(ctx context.Context, order *models.OrderCreate) (string, error)
}

type OrderServ struct {
	repo OrderStore
}

func NewOrderService(repo OrderStore) *OrderServ {
	return &OrderServ{repo: repo}
}

func (s *OrderServ) CreateOrder(ctx context.Context, order *models.OrderCreate) (string, error) {
	var items []models.Item
	var orderDetails models.Order
	id := uuid.New()
	orderDetails.ID = id
	orderDetails.UserID = order.UserID
	orderDetails.RestaurantID = order.RestaurantID
	orderDetails.TotalPrice = order.TotalPrice

	for _, item := range order.Items {
		item.ID = uuid.New()
		item.OrderID = id
		items = append(items, item)
	}
	err := s.repo.CreateOrder(ctx, &orderDetails, &items)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
