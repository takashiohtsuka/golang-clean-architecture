-- +goose Up
-- +goose StatementBegin
CREATE TABLE woman_images (
    id INT NOT NULL AUTO_INCREMENT,
    woman_id INT NOT NULL COMMENT 'woman id',
    path varchar(1000) NOT NULL COMMENT 's3 path',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    INDEX woman_image_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_woman_images_woman FOREIGN KEY (woman_id) REFERENCES women (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='woman images';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE woman_images;
-- +goose StatementEnd
