package handlers

import (
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login (c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.Queries.Login(context.Background(), request.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	token, err := GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully loged in",
		"user": gin.H{
			"id": user.ID,
			"email": user.Email,
			"username": user.Username,
		},
	})
}