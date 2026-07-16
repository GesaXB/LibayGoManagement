package middlewares

import (
	"net/http"
	"strings"

	"github.com/GesaXB/LibayGoManagement/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := strings.TrimPrefix(
			c.GetHeader("Authorization"),
			"Bearer ",
		)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
