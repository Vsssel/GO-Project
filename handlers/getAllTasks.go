package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllTasks(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userIDInt, ok := userID.(int)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }

    tasks, err := h.Queries.GetAllTasks(context.Background(), int32(userIDInt))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}
