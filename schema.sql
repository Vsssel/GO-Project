CREATE TABLE "statuses" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL
);

CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" TEXT NOT NULL
);

CREATE TABLE "tasks" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "status_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT fk_status FOREIGN KEY ("status_id") REFERENCES "statuses"("id") ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE INDEX idx_tasks_status_id ON "tasks"("status_id");
CREATE INDEX idx_tasks_user_id ON "tasks"("user_id");
CREATE INDEX idx_users_email ON "users"("email");
