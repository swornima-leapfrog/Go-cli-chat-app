-- +goose Up
-- SQL for applying the migration (creating the table)

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()  
);


-- +goose Down
-- SQL for rolling back the migration (dropping the table)

DROP TABLE IF EXISTS users;
