CREATE TABLE IF NOT EXISTS food_categories (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `tag`                   TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)',
    `title`                 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '類別名稱',
    `is_deleted`             TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否刪除',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;