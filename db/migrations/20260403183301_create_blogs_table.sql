-- +goose Up
-- +goose StatementBegin
CREATE TABLE blogs (
    id INT NOT NULL AUTO_INCREMENT,
    woman_id INT NOT NULL COMMENT 'woman id',
    title varchar(255) NOT NULL COMMENT 'title',
    body text DEFAULT NULL COMMENT 'body',
    is_published tinyint(1) NOT NULL DEFAULT 0 COMMENT 'published flag',
    created_at datetime DEFAULT NULL COMMENT 'created at',
    updated_at datetime DEFAULT NULL COMMENT 'updated at',
    deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
    INDEX blog_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_blogs_woman FOREIGN KEY (woman_id) REFERENCES women (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='blogs';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE blogs;
-- +goose StatementEnd
