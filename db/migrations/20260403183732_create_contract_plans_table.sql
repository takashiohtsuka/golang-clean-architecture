-- +goose Up
-- +goose StatementBegin
CREATE TABLE contract_plans (
    id INT NOT NULL AUTO_INCREMENT,
    code varchar(255) NOT NULL COMMENT 'contract plan code',
    INDEX contract_plan_id (id),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='contract plans';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contract_plans;
-- +goose StatementEnd
