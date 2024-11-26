package http

import "github.com/JerryJeager/JeagerEats/internal/service/categories"

type CategoryController struct {
	serv categories.CategorySv
}

func NewCategoryController(serv categories.CategorySv) *CategoryController {
	return &CategoryController{serv: serv}
}
