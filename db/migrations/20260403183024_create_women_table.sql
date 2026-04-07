-- +goose Up
-- +goose StatementBegin
CREATE TABLE women (
    id INT NOT NULL AUTO_INCREMENT,
    company_id INT NOT NULL COMMENT 'company id',
    name varchar(255) NOT NULL COMMENT 'name',
    age INT DEFAULT NULL COMMENT 'age',
    birthplace varchar(255) DEFAULT NULL COMMENT 'birthplace',
    blood_type varchar(10) DEFAULT NULL COMMENT 'blood type',
    hobby varchar(1000) DEFAULT NULL COMMENT 'hobby (comma separated)',
    is_active tinyint(1) NOT NULL DEFAULT 1 COMMENT 'active flag',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX woman_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_women_company FOREIGN KEY (company_id) REFERENCES companies (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='women';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE women;
-- +goose StatementEnd
