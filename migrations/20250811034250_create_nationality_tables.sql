-- +goose Up
-- +goose StatementBegin
CREATE TABLE nationality (
    id SERIAL PRIMARY KEY,
    nationality_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE nationality;
-- +goose StatementEnd
