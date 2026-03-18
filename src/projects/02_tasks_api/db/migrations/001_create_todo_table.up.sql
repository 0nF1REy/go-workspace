-- Migration: Create todos table
CREATE TABLE IF NOT EXISTS tasks.todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    done BOOLEAN NOT NULL DEFAULT FALSE
);
