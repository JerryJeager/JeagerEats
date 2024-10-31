package middleware

import (
	"fmt"
	"net/http"

	auth "github.com/JerryJeager/JeagerEats/internal/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, role, restaurantID, err := auth.ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":     "Bad request",
				"message":    "Authentication failed",
				"statusCode": http.StatusUnauthorized,
			})
			fmt.Println(err)
			c.Abort()
			return
		}

		c.Set("user_id", id) 
		c.Set("role", role)
		c.Set("restaurant_id", restaurantID)

		c.Next()
	}
}
