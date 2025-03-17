package main

import (
	"database/sql"
	"log"
	"to-do-list-app/handlers"
	"to-do-list-app/internal/repository"
	"to-do-list-app/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var queries *repository.Queries

func main() {
	ConnectDB()
	r := gin.Default()
	handler := &handlers.Handler{Queries: queries}

	r.GET("/tasks", middleware.RequireAuth, handler.GetAllTasks)
	r.POST("/task",middleware.RequireAuth, handler.CreateTask)
	r.PUT("/tasks/:id", middleware.RequireAuth, handler.UpdateTasksStatus)
	r.DELETE("/tasks/:id", middleware.RequireAuth, handler.DeleteTask)
	r.POST("/signup",  handler.SignupHandler)
	r.GET("/login", handler.Login)

	r.Run(":8080")
}

func ConnectDB() {
	conn := "postgres://assel:password@localhost:4444/to-do-app?sslmode=disable"
	dbconn, err := sql.Open("postgres", conn)

	// defer dbconn.Close()

	queries = repository.New(dbconn)

	if err != nil {
		log.Fatal(err)
	}
}

