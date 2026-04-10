-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores MODIFY COLUMN district_id INT NOT NULL COMMENT 'district id';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores MODIFY COLUMN district_id INT DEFAULT NULL COMMENT 'district id';
-- +goose StatementEnd
