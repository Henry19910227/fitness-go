CREATE TABLE IF NOT EXISTS cards (
    `user_id` INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `card_id` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '身分證字號',
    `front_image` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '身分證正面照',
    `back_image` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '身分證背面照',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_cards_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;