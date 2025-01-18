-- +goose Up
-- +goose StatementBegin
ALTER TABLE staffs ADD is_active ENUM('true', 'false') default 'true' NOT NULL;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
ALTER TABLE staffs DROP is_active;
-- +goose StatementEnd
