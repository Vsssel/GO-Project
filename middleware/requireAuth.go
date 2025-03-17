package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userID, exists := claims["sub"].(float64)

		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println(claims["foo"], claims["nbf"])

		c.Set("userID", userID)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}