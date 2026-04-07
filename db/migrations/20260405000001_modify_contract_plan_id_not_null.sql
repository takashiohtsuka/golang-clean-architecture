-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores MODIFY COLUMN contract_plan_id INT NOT NULL COMMENT 'contract plan id';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores MODIFY COLUMN contract_plan_id INT DEFAULT NULL COMMENT 'contract plan id';
-- +goose StatementEnd
