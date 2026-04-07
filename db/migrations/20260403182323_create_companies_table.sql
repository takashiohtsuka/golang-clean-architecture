-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies (
    id INT NOT NULL AUTO_INCREMENT,
    business_type_id INT NOT NULL COMMENT 'business type id',
    name varchar(255) NOT NULL COMMENT 'company name',
    is_active tinyint(1) NOT NULL DEFAULT 1 COMMENT 'active flag',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX company_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_companies_business_type FOREIGN KEY (business_type_id) REFERENCES business_types (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='companies';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE companies;
-- +goose StatementEnd
