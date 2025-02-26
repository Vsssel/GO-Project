package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"to-do-list-app/handlers"
	"to-do-list-app/internal/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var queries *repository.Queries

func main() {
	ConnectDB()
	fmt.Println(queries.GetAllTasks(context.Background()))
	r := gin.Default()
	handler := &handlers.Handler{Queries: queries}

	r.GET("/tasks", handler.GetAllTasks)
	r.POST("/task", handler.CreateTask)
	r.PUT("/tasks/:id", handler.UpdateTasksStatus)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	r.Run(":8080")
}

func ConnectDB() {
	conn := "postgres://assel:password@localhost:5431/to_do_list?sslmode=disable"
	dbconn, err := sql.Open("postgres", conn)

	// defer dbconn.Close()

	queries = repository.New(dbconn)

	if err != nil {
		log.Fatal(err)
	}
}

