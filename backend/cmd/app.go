package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/JeagerEats/manualwire"
	"github.com/JerryJeager/JeagerEats/middleware"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()
var restaurantController = manualwire.GetRestaurantController()
var menuController = manualwire.GetMenuController()

func ExecuteApiRoutes() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to JeagerEats",
		})
	})

	api := router.Group("/api/v1")
	users := api.Group("/users")
	restaurants := api.Group("/restaurants")
	menus := restaurants.Group("menus")

	users.POST("/signup", userController.CreateUser)
	users.POST("/login", userController.Login)

	restaurants.PATCH("/profile", middleware.JwtAuthMiddleware(), restaurantController.UpdateRestaurant)
	restaurants.PATCH("/profile/img", middleware.JwtAuthMiddleware(), middleware.FileUploadMiddleware(), restaurantController.UpdateRestaurantProfileImg)
	restaurants.PATCH("/active", middleware.JwtAuthMiddleware(), restaurantController.UpdateRestaurantIsActive)
	restaurants.GET("/:id", restaurantController.GetRestaurantPublicProfile)
	restaurants.GET("", restaurantController.GetAllRestaurantPublicProfile)

	menus.POST("", middleware.JwtAuthMiddleware(), menuController.CreateMenu)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
