-- Criação de schemas
CREATE SCHEMA IF NOT EXISTS products;

CREATE SCHEMA IF NOT EXISTS tasks;

-- Tabela da API de tasks
CREATE TABLE IF NOT EXISTS tasks.todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela da API de products
CREATE TABLE IF NOT EXISTS products.products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC
);
