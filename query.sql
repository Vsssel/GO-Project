-- name: GetAllTasks :many

SELECT * FROM tasks
WHERE user_id = $1;


-- name: CreateTask :exec

INSERT INTO tasks (title, status_id, user_id)
VALUES ($1, 1, $2);

-- name: UpdateTasksStatus :exec

UPDATE tasks
SET status_id = $1
WHERE id = $2;

-- name: DeleteTask :exec

DELETE FROM tasks WHERE id = $1;


-- name: Signup :exec 

INSERT INTO users (username, email, password)
VALUES ($1, $2, $3);


-- name: Login :one

SELECT * FROM users 
WHERE email = $1;
