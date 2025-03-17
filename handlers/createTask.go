package handlers

import (
	"context"
	"net/http"
	"to-do-list-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var reqBody struct {
		title string
	}

	userID, exists := c.Get("user_id")

	if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userIDInt, ok := userID.(int32)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err := h.Queries.CreateTask(context.Background(), repository.CreateTaskParams{
		Title: reqBody.title,
		UserID: userIDInt,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Task Created Successfully")

}