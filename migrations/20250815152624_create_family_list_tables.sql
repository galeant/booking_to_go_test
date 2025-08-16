-- +goose Up
-- +goose StatementBegin
CREATE TABLE family_list (
    fl_id SERIAL PRIMARY KEY,
    cst_id INT NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_family_list_user 
        FOREIGN KEY (cst_id) 
        REFERENCES customer(cst_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE family_list;
-- +goose StatementEnd
