-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS log_table (
                                         id SERIAL PRIMARY KEY,
                                         timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         action TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS log_table;
-- +goose StatementEnd
