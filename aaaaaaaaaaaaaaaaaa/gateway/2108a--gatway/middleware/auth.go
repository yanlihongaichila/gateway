package middleware

import (
	"gateway/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	userID, err := utils.CheckToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
	}
	c.Set("user_id", userID)
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	},
	)
}
