-- +goose Up
-- +goose StatementBegin
CREATE TABLE photos (
    id INT NOT NULL AUTO_INCREMENT,
    blog_id INT NOT NULL COMMENT 'blog id',
    url varchar(1000) NOT NULL COMMENT 'photo url',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    INDEX photo_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_photos_blog FOREIGN KEY (blog_id) REFERENCES blogs (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='photos';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE photos;
-- +goose StatementEnd
