package middleware

import (
	"fmt"
	"getway/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	fmt.Println("********token")
	fmt.Println(token)

	userID, err := utils.CheckToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
	}
	c.Set("user_id", userID)
}
