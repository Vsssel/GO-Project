package handlers

import (
	"context"
	"net/http"
	"strconv"
	"to-do-list-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateTasksStatus (c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.Queries.UpdateTasksStatus(context.Background(), repository.UpdateTasksStatusParams{StatusID: int32(request.StatusId), ID: int32(id)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Status Updated Successfully")
}

type Request struct {
	StatusId int `json:"statusId"`
}