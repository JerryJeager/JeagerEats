package models

import (
	// "fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uuid.UUID  `json:"id"`
	Email       string     `json:"email" binding:"required" gorm:"unique"`
	Password    string     `json:"password" binding:"required"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Role        string     `json:"role" binding:"required"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const (
	CUSTOMER = "customer"
	VENDOR   = "vendor"
	RIDER    = "rider"
)

func IsValidRole(role string) bool {
	return role == CUSTOMER || role == VENDOR || role == RIDER
}

func (user *User) HashPassword() error {
	user.Password = html.EscapeString(strings.TrimSpace(user.Password))
	log.Print(user.Email, user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	// fmt.log("password: ", password)
	log.Print(password)
	password = html.EscapeString(strings.TrimSpace(password))
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
