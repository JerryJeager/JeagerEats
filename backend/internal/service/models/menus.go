package models

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID           uuid.UUID `json:"id"`
	RestaurantID uuid.UUID `json:"restaurant_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	IsAvailable  bool      `json:"is_available"`
	ImgURL       string    `json:"img_url"`
	Stock        int       `json:"stock"`
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
