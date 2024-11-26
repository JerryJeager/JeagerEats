package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/categories"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	serv categories.CategorySv
}

func NewCategoryController(serv categories.CategorySv) *CategoryController {
	return &CategoryController{serv: serv}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorJson(err, "category must include name and description fields"))
		return
	}

	id, err := c.serv.CreateCategory(ctx, &category)
	if err != nil {
		ctx.JSON(http.StatusConflict, GetErrorJson(err, ""))
	}

	ctx.JSON(http.StatusCreated, id)
}
