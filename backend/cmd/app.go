package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/JeagerEats/manualwire"
	"github.com/JerryJeager/JeagerEats/middleware"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to JeagerEats",
		})
	})

	userController := manualwire.GetUserController()
	restaurantController := manualwire.GetRestaurantController()
	menuController := manualwire.GetMenuController()
	riderController := manualwire.GetRiderController()
	category := manualwire.GetCategoryController()
	orderController := manualwire.GetOrderController()

	api := router.Group("/api/v1")
	users := api.Group("/users")
	restaurants := api.Group("/restaurants")
	menus := api.Group("menus")
	riders := api.Group("riders")
	categories := api.Group("categories")
	orders := api.Group("orders")

	users.POST("/signup", userController.CreateUser)
	users.POST("/login", userController.Login)
	users.GET("", middleware.JwtAuthMiddleware(), userController.GetUser)

	restaurants.PATCH("/profile", middleware.JwtAuthMiddleware(), restaurantController.UpdateRestaurant)
	restaurants.PATCH("/profile/img", middleware.JwtAuthMiddleware(), middleware.FileUploadMiddleware(), restaurantController.UpdateRestaurantProfileImg)
	restaurants.PATCH("/active", middleware.JwtAuthMiddleware(), restaurantController.UpdateRestaurantIsActive)
	restaurants.GET("/:id", restaurantController.GetRestaurantPublicProfile)
	restaurants.GET("", restaurantController.GetAllRestaurantPublicProfile)
	restaurants.GET("/self", middleware.JwtAuthMiddleware(), restaurantController.GetRestaurant)

	menus.POST("", middleware.JwtAuthMiddleware(), menuController.CreateMenu)
	menus.PATCH("/img/:id", middleware.JwtAuthMiddleware(), middleware.FileUploadMiddleware(), menuController.UpdateMenuImage)
	menus.GET("", menuController.GetMenus)
	menus.GET("/restaurants/:id", menuController.GetMenusByRestaurantID)
	menus.GET("/:id", menuController.GetMenuByID)
	menus.DELETE("/:id", middleware.JwtAuthMiddleware(), menuController.DeleteMenu)
	menus.PATCH("/:id", middleware.JwtAuthMiddleware(), menuController.UpdateMenu)

	riders.PATCH("", middleware.JwtAuthMiddleware(), riderController.UpdateRider)
	riders.PATCH("/profile/img", middleware.JwtAuthMiddleware(), middleware.FileUploadMiddleware(), riderController.UpdateRiderProfileImg)

	categories.POST("", category.CreateCategory)

	orders.POST("", middleware.JwtAuthMiddleware(), orderController.CreateOrder)
	orders.POST("/:id/accept", orderController.AcceptOrder)
	orders.PUT("/:id/status", middleware.JwtAuthMiddleware(), orderController.UdpateOrderStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
