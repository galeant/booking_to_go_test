-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer (
    cst_id SERIAL PRIMARY KEY,
    nationality_id INT NOT NULL,
    cst_name VARCHAR(50) NOT NULL,
    cst_dob DATE NOT NULL,
    cst_phoneNum VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_user_nationality FOREIGN KEY (nationality_id) REFERENCES users(cst_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer;
-- +goose StatementEnd
