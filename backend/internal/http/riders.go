package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/riders"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RiderController struct {
	serv riders.RiderSv
}

func NewRiderController(serv riders.RiderSv) *RiderController {
	return &RiderController{serv: serv}
}

func (r *RiderController) UpdateRider(ctx *gin.Context) {
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

	var rider models.RiderUpdate
	if err := ctx.ShouldBindJSON(&rider); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.serv.UpdateRider(ctx, userID, &rider); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Rider updated successfully"})
}
