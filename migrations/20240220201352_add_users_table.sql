-- +goose Up
-- +goose StatementBegin
-- DEFAULT USERS
INSERT INTO users (user_name, email, password_hash, role_id)
VALUES
    ('admin', 'admin@chat', '827ccb0eea8a706c4c34a16891f84e7b', 2),
    ('user', 'user@chat', '827ccb0eea8a706c4c34a16891f84e7b', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE users;
-- +goose StatementEnd
