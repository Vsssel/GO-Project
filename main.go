package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	queries = repository.New(dbconn)

	if err != nil {
		log.Fatal(err)
	}
}

