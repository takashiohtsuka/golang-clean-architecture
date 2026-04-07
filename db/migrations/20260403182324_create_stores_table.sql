-- +goose Up
-- +goose StatementBegin
CREATE TABLE stores (
    id INT NOT NULL AUTO_INCREMENT,
    company_id INT NOT NULL COMMENT 'company id',
    name varchar(255) NOT NULL COMMENT 'store name',
    is_active tinyint(1) NOT NULL DEFAULT 1 COMMENT 'active flag',
    open_status varchar(20) NOT NULL DEFAULT 'private' COMMENT 'open status: public / restricted / private',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX store_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_stores_company FOREIGN KEY (company_id) REFERENCES companies (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='stores';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stores;
-- +goose StatementEnd
