-- +goose Up
-- +goose StatementBegin
CREATE TABLE nationality (
    nationality_id SERIAL PRIMARY KEY,
    nationality_name VARCHAR(50) NOT NULL,
    nationality_code VARCHAR(2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE nationality;
-- +goose StatementEnd
