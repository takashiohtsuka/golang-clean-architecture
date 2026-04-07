-- +goose Up
-- +goose StatementBegin
ALTER TABLE companies
    DROP FOREIGN KEY fk_companies_business_type,
    DROP COLUMN business_type_id;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE stores
    ADD COLUMN business_type_id INT NOT NULL COMMENT 'business type id',
    ADD CONSTRAINT fk_stores_business_type FOREIGN KEY (business_type_id) REFERENCES business_types (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores
    DROP FOREIGN KEY fk_stores_business_type,
    DROP COLUMN business_type_id;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE companies
    ADD COLUMN business_type_id INT NOT NULL COMMENT 'business type id',
    ADD CONSTRAINT fk_companies_business_type FOREIGN KEY (business_type_id) REFERENCES business_types (id);
-- +goose StatementEnd
