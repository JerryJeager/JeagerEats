package http

import (
	"mime/multipart"
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/restaurants"
	"github.com/JerryJeager/JeagerEats/internal/utils"
	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	serv restaurants.RestaurantSv
}

func NewRestaurantController(serv restaurants.RestaurantSv) *RestaurantController {
	return &RestaurantController{serv: serv}
}

func (c *RestaurantController) UpdateRestaurant(ctx *gin.Context) {
	vendor, err := GetVendor(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var restaurant models.RestaurantUpdate
	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.serv.UpdateRestaurant(ctx, vendor.UserID, &restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant updated successfully"})
}

func (c *RestaurantController) UpdateRestaurantProfileImg(ctx *gin.Context) {
	vendor, err := GetVendor(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	filename, ok := ctx.Get("filePath")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
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

	err = c.serv.UpdateRestaurantProfileImg(ctx, vendor.UserID, imageUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant profile image updated successfully", "image_url": imageUrl})
}
