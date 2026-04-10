-- +goose Up
-- +goose StatementBegin
CREATE TABLE prefectures (
    id        INT NOT NULL AUTO_INCREMENT,
    region_id INT NOT NULL COMMENT 'region id',
    name      VARCHAR(50) NOT NULL COMMENT '都道府県名',
    INDEX prefecture_id (id),
    PRIMARY KEY(id),
    CONSTRAINT fk_prefectures_region FOREIGN KEY (region_id) REFERENCES regions (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='都道府県';
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO prefectures (region_id, name) VALUES
    -- 北海道
    (1, '北海道'),
    -- 東北
    (2, '青森県'),
    (2, '岩手県'),
    (2, '宮城県'),
    (2, '秋田県'),
    (2, '山形県'),
    (2, '福島県'),
    -- 関東
    (3, '茨城県'),
    (3, '栃木県'),
    (3, '群馬県'),
    (3, '埼玉県'),
    (3, '千葉県'),
    (3, '東京都'),
    (3, '神奈川県'),
    -- 中部
    (4, '新潟県'),
    (4, '富山県'),
    (4, '石川県'),
    (4, '福井県'),
    (4, '山梨県'),
    (4, '長野県'),
    (4, '岐阜県'),
    (4, '静岡県'),
    (4, '愛知県'),
    -- 近畿
    (5, '三重県'),
    (5, '滋賀県'),
    (5, '京都府'),
    (5, '大阪府'),
    (5, '兵庫県'),
    (5, '奈良県'),
    (5, '和歌山県'),
    -- 中国
    (6, '鳥取県'),
    (6, '島根県'),
    (6, '岡山県'),
    (6, '広島県'),
    (6, '山口県'),
    -- 四国
    (7, '徳島県'),
    (7, '香川県'),
    (7, '愛媛県'),
    (7, '高知県'),
    -- 九州・沖縄
    (8, '福岡県'),
    (8, '佐賀県'),
    (8, '長崎県'),
    (8, '熊本県'),
    (8, '大分県'),
    (8, '宮崎県'),
    (8, '鹿児島県'),
    (8, '沖縄県');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE prefectures;
-- +goose StatementEnd
