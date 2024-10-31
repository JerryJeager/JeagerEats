package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/menus"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/gin-gonic/gin"
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
	if err := ctx.ShouldBindJSON(&menu); err != nil{
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
