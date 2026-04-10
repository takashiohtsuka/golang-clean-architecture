-- +goose Up
-- +goose StatementBegin
CREATE TABLE regions (
    id   INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL COMMENT '地方名',
    INDEX region_id (id),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='地方';
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO regions (name) VALUES
    ('北海道'),
    ('東北'),
    ('関東'),
    ('中部'),
    ('近畿'),
    ('中国'),
    ('四国'),
    ('九州・沖縄');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE regions;
-- +goose StatementEnd
