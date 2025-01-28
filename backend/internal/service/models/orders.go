package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	RestaurantID    uuid.UUID  `json:"restaurant_id"`
	RiderID         *uuid.UUID `json:"rider_id"`
	RefID           string     `json:"ref_id"`
	Status          string     `json:"status"`
	DeliveryAddress string     `json:"delivery_address"`
	DeliveryFee     int        `json:"delivery_fee"`
	TotalPrice      float64    `json:"total_price"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
type OrderCreate struct {
	ID              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	RestaurantID    uuid.UUID  `json:"restaurant_id" binding:"required"`
	RiderID         *uuid.UUID `json:"rider_id"`
	Status          string     `json:"status"`
	TotalPrice      float64    `json:"total_price" binding:"required"`
	DeliveryAddress string     `json:"delivery_address"`
	DeliveryFee     int        `json:"delivery_fee"`
	Items           []Item     `json:"items" binding:"required"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

type Item struct {
	ID           uuid.UUID  `json:"id"`
	OrderID      uuid.UUID  `json:"order_id"`
	MenuID       uuid.UUID  `json:"menu_id" binding:"required"`
	Quantity     int        `json:"quantity" binding:"required"`
	PricePerItem float64    `json:"price_per_item" binding:"required"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type MenuOrderSummary struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Status map[string]bool

var Statuses Status = map[string]bool{
	"pending":    true,
	"accepted":   true,
	"in_transit": true,
	"delivered":  true,
}

const (
	PENDING   = "pending"
	ACCEPTED  = "accepted"
	INTRANSIT = "in_transit"
	DELIVERED = "delivered"
)

type OrderRiderUpdate struct {
	RiderID uuid.UUID `json:"rider_id" binding:"required"`
}

type OrderStatusUpdate struct {
	Status string `json:"status" binding:"required"`
}
