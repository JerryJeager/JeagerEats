package http

import (
	"mime/multipart"
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/riders"
	"github.com/JerryJeager/JeagerEats/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RiderController struct {
	serv riders.RiderSv
}

func NewRiderController(serv riders.RiderSv) *RiderController {
	return &RiderController{serv: serv}
}

func (c *RiderController) UpdateRider(ctx *gin.Context) {
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

	if err := c.serv.UpdateRider(ctx, userID, &rider); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Rider updated successfully"})
}

func (c *RiderController) UpdateRiderProfileImg(ctx *gin.Context) {
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

	filename, ok := ctx.Get("filePath")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
		return
	}

	file, ok := ctx.Get("file")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	imageUrl, err := utils.UploadToCloudinary(file.(multipart.File), filename.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.serv.UpdateRiderProfileImg(ctx, userID, imageUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Rider profile image updated successfully", "image_url": imageUrl})
}
