CREATE TABLE IF NOT EXISTS rdas (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `tdee`                  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'TDEE',
    `calorie`               INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '目標熱量',
    `protein`               INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '蛋白質(克)',
    `fat`                   INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '脂肪(克)',
    `carbs`                 INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '碳水化合物(克)',
    `grain`                 INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '穀物類(份)',
    `vegetable`             INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '蔬菜類(份)',
    `fruit`                 INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '水果類(份)',
    `meat`                  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '蛋豆魚肉類(份)',
    `dairy`                 INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '乳製品類(份)',
    `nut`                   INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '堅果類',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_rdas_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;