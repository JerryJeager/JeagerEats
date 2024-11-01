package http

import (
	"mime/multipart"
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/menus"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MenuController struct {
	serv menus.MenuSv
}

func NewMenuController(serv menus.MenuSv) *MenuController {
	return &MenuController{serv: serv}
}

func (c *MenuController) CreateMenu(ctx *gin.Context) {
	vendor, err := GetVendor(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var menu models.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := c.serv.CreateMenu(ctx, *vendor.RestaurantID, &menu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (c *MenuController) UpdateMenuImage(ctx *gin.Context) {
	var menuIDPP MenuIDPathParam
	if err := ctx.ShouldBindUri(&menuIDPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menuID, err := uuid.Parse(menuIDPP.ID)
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
	if err := c.serv.UpdateMenuImage(ctx, menuID, imageUrl); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Menu image updated successfully", "image_url": imageUrl})
}

func (c *MenuController) GetMenusByRestaurantID(ctx *gin.Context) {
	var restaurantIDPP RestaurantIDPathParam
	if err := ctx.ShouldBindUri(&restaurantIDPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	restaurantID, err := uuid.Parse(restaurantIDPP.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menus, err := c.serv.GetMenusByRestaurantID(ctx, restaurantID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menus)
}

func (c *MenuController) GetMenuByID(ctx *gin.Context) {
	var menuIDPP MenuIDPathParam
	if err := ctx.ShouldBindUri(&menuIDPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menuID, err := uuid.Parse(menuIDPP.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menu, err := c.serv.GetMenuByID(ctx, menuID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *menu)
}

func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := c.serv.GetMenus(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menus)
}

func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	var menuIDPP MenuIDPathParam
	if err := ctx.ShouldBindUri(&menuIDPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menuID, err := uuid.Parse(menuIDPP.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.serv.DeleteMenu(ctx, menuID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	var menuIDPP MenuIDPathParam
	if err := ctx.ShouldBindUri(&menuIDPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menuID, err := uuid.Parse(menuIDPP.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var menuUpdate models.MenuUpdate
	if err := ctx.ShouldBindJSON(&menuUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.serv.UpdateMenu(ctx, menuID, &menuUpdate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully"})
}
