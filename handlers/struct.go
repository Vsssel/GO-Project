package handlers

import "to-do-list-app/internal/repository"

type Handler struct {
	Queries *repository.Queries
}