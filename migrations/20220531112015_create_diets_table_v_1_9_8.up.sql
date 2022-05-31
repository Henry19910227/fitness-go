CREATE TABLE IF NOT EXISTS diets (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `rda_id`                INT(11) UNSIGNED COMMENT '建議營養id',
    `schedule_at`           DATETIME NOT NULL DEFAULT NOW() COMMENT '排程日期',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_diets` (`user_id`, `schedule_at`),
    CONSTRAINT fk_diets_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_diets_rda_id_to_rdas_id FOREIGN KEY (rda_id) REFERENCES rdas(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;