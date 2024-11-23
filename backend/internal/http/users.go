package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	id, token, role, err := c.serv.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "id": id, "token": token, "role": role})
}

func (c *UserController) GetUser(ctx *gin.Context) {
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

	user, err := c.serv.GetUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, *user)
}
