package models

import "time"

import "github.com/google/uuid"


type Rider struct {
	ID               uuid.UUID  `json:"id"`
	UserID           uuid.UUID  `json:"user_id"`
	VehicleType      string     `json:"vehicle_type"`
	LicenseNumber    string     `json:"license_number"`
	IsActive         bool       `json:"is_active"`
	ProfileImg       string     `json:"profile_img"`
	OpeningTime      time.Time  `json:"opening_time"`
	ClosingTime      time.Time  `json:"closing_time"`
	CurrentLatitude  float64    `json:"current_latitude"`
	CurrentLongitude float64    `json:"current_longitude"`
	Rating           int        `json:"rating"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

type RiderUpdate struct {
	VehicleType      string     `json:"vehicle_type"`
	LicenseNumber    string     `json:"license_number"`
	OpeningTime      time.Time  `json:"opening_time"`
	ClosingTime      time.Time  `json:"closing_time"`
	CurrentLatitude  float64    `json:"current_latitude"`
	CurrentLongitude float64    `json:"current_longitude"`
	IsActive         bool       `json:"is_active"`
}


type Riders []Rider