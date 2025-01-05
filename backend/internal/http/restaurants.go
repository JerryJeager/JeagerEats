package http

import (
	"mime/multipart"
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/restaurants"
	"github.com/JerryJeager/JeagerEats/internal/utils"
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

func (c *RestaurantController) UpdateRestaurantIsActive(ctx *gin.Context) {
	vendor, err := GetVendor(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var isActive models.IsActive
	if err := ctx.ShouldBindJSON(&isActive); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.serv.UpdateRestaurantIsActive(ctx, vendor.UserID, &isActive)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant active status updated successfully"})
}

func (c *RestaurantController) GetRestaurantPublicProfile(ctx *gin.Context) {
	var restaurantIDPathParam RestaurantIDPathParam
	if err := ctx.ShouldBindUri(&restaurantIDPathParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(restaurantIDPathParam.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	restaurant, err := c.serv.GetRestaurantPublicProfile(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *restaurant)
}

func (c *RestaurantController) GetAllRestaurantPublicProfile(ctx *gin.Context) {
	restaurants, err := c.serv.GetAllRestaurantPublicProfile(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *restaurants)
}

func (c *RestaurantController) GetRestaurant(ctx *gin.Context){
	vendor, err := GetVendor(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	restaurant, err := c.serv.GetRestaurant(ctx, *&vendor.UserID)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, restaurant)
}