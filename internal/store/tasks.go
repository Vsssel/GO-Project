package store

import (
	"context"
	"database/sql"
)

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
	UserID int `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TasksStore struct {
	db *sql.DB
}

func (s *TasksStore) Create(ctx context.Context, task *Task) error {
	query := `
		INSERT INTO tasks (title, status, user_id)
		VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx, 
		query, 
		task.Title, 
		task.Status, 
		task.UserID,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}