// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

type Status struct {
	ID    int32
	Title string
}

type Task struct {
	ID       int32
	Title    string
	StatusID int32
}
