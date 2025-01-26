package http

import (
	"errors"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetVendor(ctx *gin.Context) (*models.Vendor, error) {
	var restaurantID uuid.UUID
	userIDCtx, err := GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		return nil, err
	}

	role, err := GetRole(ctx)
	if role != "vendor" {
		return nil, errors.New("unauthorized")
	}

	restaurantIDCtx, err := GetRestaurantID(ctx)
	if err != nil {
		return nil, err
	}
	restaurantID, err = uuid.Parse(restaurantIDCtx)
	if err != nil {
		return nil, err
	}

	return &models.Vendor{UserID: userID, Role: role, RestaurantID: &restaurantID}, nil
}
func GetUser(ctx *gin.Context) (*models.Vendor, error) {
	var restaurantID uuid.UUID
	userIDCtx, err := GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		return nil, err
	}

	role, err := GetRole(ctx)
	if role != "customer" {
		return nil, errors.New("unauthorized")
	}

	return &models.Vendor{UserID: userID, Role: role, RestaurantID: &restaurantID}, nil
}
