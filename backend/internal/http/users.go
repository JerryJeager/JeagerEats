package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := c.serv.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "user with email already exists"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "id": id})
}

func (c *UserController) Login(ctx *gin.Context) {
	var user models.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, token, err := c.serv.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Login failed", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "id": id, "token": token})
}
