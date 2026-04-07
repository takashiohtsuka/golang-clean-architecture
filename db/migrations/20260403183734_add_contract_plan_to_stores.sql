-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores ADD COLUMN contract_plan_id INT DEFAULT NULL COMMENT 'contract plan id',
    ADD CONSTRAINT fk_stores_contract_plan FOREIGN KEY (contract_plan_id) REFERENCES contract_plans (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores DROP FOREIGN KEY fk_stores_contract_plan,
    DROP COLUMN contract_plan_id;
-- +goose StatementEnd
