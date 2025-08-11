-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_rent_books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    return_date DATETIME NOT NULL,
    status ENUM('borrowed', 'returned') NOT NULL DEFAULT 'borrowed',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_rent_books_users FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_user_rent_books_books FOREIGN KEY (book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_rent_books;
-- +goose StatementEnd
