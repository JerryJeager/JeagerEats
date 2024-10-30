package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/restaurants"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RestaurantController struct {
	serv restaurants.RestaurantSv
}

func NewRestaurantController(serv restaurants.RestaurantSv) *RestaurantController {
	return &RestaurantController{serv: serv}
}

func (c *RestaurantController) UpdateRestaurant(ctx *gin.Context) {
	userIDCtx, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role, err := GetRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var restaurant models.RestaurantUpdate
	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if role != "vendor" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
		return
	}

	err = c.serv.UpdateRestaurant(ctx, userID, &restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant updated successfully"})
}
