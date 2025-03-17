package handlers

import (
	"context"
	"net/http"
	"to-do-list-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var reqBody struct {
		Title string `json:"title"` // Capitalized to make it exported
	}

	// Retrieve user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDInt, ok := userID.(int32)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save task in DB
	err := h.Queries.CreateTask(context.Background(), repository.CreateTaskParams{
		Title:  reqBody.Title,
		UserID: userIDInt,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task Created Successfully"})
}
