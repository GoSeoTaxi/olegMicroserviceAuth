-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS roles (
                                     id SERIAL PRIMARY KEY,
                                     role_name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     user_name VARCHAR(50) UNIQUE NOT NULL,
                                     email VARCHAR(100) UNIQUE NOT NULL,
                                     password_hash VARCHAR(32) NOT NULL,
                                     role_id INT,
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                     FOREIGN KEY (role_id) REFERENCES roles(id)
);

INSERT INTO roles (role_name) VALUES ('USER'),('ADMIN');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd