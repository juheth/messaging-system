package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token in header"})
			c.Abort()
			return
		}

		fmt.Println("Token recibido:", tokenString)

		// Eliminar el prefijo "Bearer " del token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			fmt.Printf("Error al validar el token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Printf("Token válido. Reclamaciones: %+v\n", claims)

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
