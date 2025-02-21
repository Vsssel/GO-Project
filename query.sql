-- name: GetAllTasks :many

SELECT * FROM tasks;

-- name: GetAllCompletedTasks :many

SELECT * FROM tasks 
WHERE status_id = 2;

-- name: GetAllCreatedTasks :many

SELECT * FROM tasks 
WHERE status_id = 1;

-- name: CreateTask :exec

INSERT INTO tasks (title, status_id)
VALUES ($1, 1);

-- name: UpdateTasksStatus :exec

UPDATE tasks
SET status_id = $1
WHERE id = $2;
