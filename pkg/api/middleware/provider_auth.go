package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ProviderAuthMiddleware(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MissingAuthorizationToken})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte("providersecret"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthorizationToken})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthorizationToken})
		c.Abort()
		return
	}

	fmt.Println("role", claims)

	role, ok := claims["role"].(string)
	if !ok || role != "provider" {
		c.JSON(http.StatusForbidden, gin.H{"error": UnAuthorizedAccess})
		c.Abort()
		return
	}

	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": TokenProblem})
		c.Abort()
		return
	}

	userid := int(id)
	c.Set("Uid", userid)

	// c.Next()
}
