-- +goose Up
-- +goose StatementBegin
CREATE TABLE woman_store_assignments (
    id INT NOT NULL AUTO_INCREMENT,
    woman_id INT NOT NULL COMMENT 'woman id',
    store_id INT NOT NULL COMMENT 'store id',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    INDEX woman_store_assignment_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_woman_store_assignments_woman FOREIGN KEY (woman_id) REFERENCES women (id),
    CONSTRAINT fk_woman_store_assignments_store FOREIGN KEY (store_id) REFERENCES stores (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='woman store assignments';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE woman_store_assignments;
-- +goose StatementEnd
