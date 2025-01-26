package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	RestaurantID uuid.UUID  `json:"restaurant_id"`
	RiderID      *uuid.UUID `json:"rider_id"`
	Status       string     `json:"status"`
	TotalPrice   float64    `json:"total_price"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
type OrderCreate struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	RestaurantID uuid.UUID  `json:"restaurant_id" binding:"required"`
	RiderID      *uuid.UUID `json:"rider_id"`
	Status       string     `json:"status"`
	TotalPrice   float64    `json:"total_price" binding:"required"`
	Items        []Item     `json:"items" binding:"required"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
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
