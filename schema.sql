CREATE TABLE "statuses" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL
);

CREATE TABLE "tasks" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "status_id" INTEGER NOT NULL,
    CONSTRAINT fk_status FOREIGN KEY ("status_id") REFERENCES "statuses"("id") ON DELETE CASCADE
);
