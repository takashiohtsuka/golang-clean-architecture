-- +goose Up
-- +goose StatementBegin
CREATE TABLE business_types (
    id INT NOT NULL AUTO_INCREMENT,
    code varchar(255) NOT NULL COMMENT 'business type code',
    INDEX business_type_id (id),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='business types';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE business_types;
-- +goose StatementEnd
