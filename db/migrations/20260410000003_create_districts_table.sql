-- +goose Up
-- +goose StatementBegin
CREATE TABLE districts (
    id            INT NOT NULL AUTO_INCREMENT,
    prefecture_id INT NOT NULL COMMENT 'prefecture id',
    name          VARCHAR(100) NOT NULL COMMENT '繁華街名',
    INDEX district_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_districts_prefecture FOREIGN KEY (prefecture_id) REFERENCES prefectures (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='繁華街';
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO districts (prefecture_id, name) VALUES
    -- 東京都 (prefecture_id=13)
    (13, '新宿'),
    (13, '渋谷'),
    (13, '銀座'),
    (13, '池袋'),
    (13, '上野'),
    (13, '六本木'),
    -- 神奈川県 (prefecture_id=14)
    (14, '横浜'),
    (14, '川崎'),
    -- 大阪府 (prefecture_id=27)
    (27, '梅田'),
    (27, '難波'),
    (27, '心斎橋'),
    -- 兵庫県 (prefecture_id=28)
    (28, '神戸三宮'),
    -- 愛知県 (prefecture_id=23)
    (23, '名古屋栄'),
    (23, '名駅'),
    -- 福岡県 (prefecture_id=40)
    (40, '中洲'),
    (40, '天神'),
    -- 北海道 (prefecture_id=1)
    (1, 'すすきの');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE districts;
-- +goose StatementEnd
