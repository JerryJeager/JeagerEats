package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

const (
	CUSTOMER = "customer" 
	VENDOR = "vendor"
	RIDER = "rider"
)

func IsValidRole(role string) bool{
	return role == CUSTOMER || role == VENDOR || role == RIDER
}