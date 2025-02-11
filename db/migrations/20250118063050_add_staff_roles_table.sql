-- +goose Up
-- +goose StatementBegin
CREATE TABLE staff_roles (
    id INT NOT NULL AUTO_INCREMENT,
    staff_id INT NOT NULL COMMENT 'staff id',
    role_id INT NOT NULL COMMENT 'role_id',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX idx_staff_id (staff_id),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='staff_roles';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE staff_roles;
-- +goose StatementEnd
