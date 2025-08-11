-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password LONGTEXT NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    type ENUM('admin', 'member') NOT NULL DEFAULT 'member',
    last_login TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
