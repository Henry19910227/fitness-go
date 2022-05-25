CREATE TABLE IF NOT EXISTS user_income_monthly_statistics (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `income`                INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '收益',
    `year`                  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '年份',
    `month`                 INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '月份',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_user_income_monthly_statistics` (`user_id`, `year`, `month`),
    CONSTRAINT fk_user_income_monthly_statistics_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;