-- +goose Up
-- +goose StatementBegin
CREATE TABLE management_staffs (
    id INT NOT NULL AUTO_INCREMENT,
    company_id INT NOT NULL COMMENT 'company id',
    store_id INT NOT NULL COMMENT 'store id',
    name varchar(255) NOT NULL COMMENT 'name',
    email varchar(255) NOT NULL COMMENT 'email',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX management_staff_id (id),
    UNIQUE KEY uq_management_staffs_email (email),
    PRIMARY KEY(id),
    CONSTRAINT fk_management_staffs_company FOREIGN KEY (company_id) REFERENCES companies (id),
    CONSTRAINT fk_management_staffs_store FOREIGN KEY (store_id) REFERENCES stores (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='management staffs';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE management_staffs;
-- +goose StatementEnd
