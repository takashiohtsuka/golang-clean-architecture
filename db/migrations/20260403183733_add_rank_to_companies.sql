-- +goose Up
-- +goose StatementBegin
ALTER TABLE companies ADD COLUMN `rank` varchar(255) DEFAULT NULL COMMENT 'rank';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE companies DROP COLUMN `rank`;
-- +goose StatementEnd
