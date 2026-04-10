-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores ADD COLUMN district_id INT DEFAULT NULL COMMENT 'district id' AFTER company_id,
    ADD CONSTRAINT fk_stores_district FOREIGN KEY (district_id) REFERENCES districts (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores DROP FOREIGN KEY fk_stores_district,
    DROP COLUMN district_id;
-- +goose StatementEnd
