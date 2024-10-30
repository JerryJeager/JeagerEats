package models

import "time"

import "github.com/google/uuid"

type Restaurant struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	OpeningTime time.Time  `json:"opening_time"`
	ClosingTime time.Time  `json:"closing_time"`
	CuisineType string     `json:"cuisine_type"`
	IsActive    bool       `json:"is_active"`
	ProfileImg  string     `json:"profile_img"`
	Rating      int        `json:"rating"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type RestaurantUpdate struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
	CuisineType string    `json:"cuisine_type"`
}

type Vendor struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}
