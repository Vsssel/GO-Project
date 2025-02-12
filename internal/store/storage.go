package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks Tasks
	Users Users
}

type Tasks interface {
	Create(context.Context, *Task) error
}

type Users interface {
	Create(context.Context, *User) error
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TasksStore{db},
		Users: &UsersStore{db},
	}
}