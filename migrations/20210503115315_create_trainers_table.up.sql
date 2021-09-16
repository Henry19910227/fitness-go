CREATE TABLE IF NOT EXISTS trainers (
    `user_id` INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '教練本名',
    `nickname` VARCHAR(20) NOT NULL UNIQUE DEFAULT '' COMMENT '教練暱稱',
    `avatar` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '教練大頭照',
    `trainer_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '教練狀態(1:正常/2:審核中/3:停權)',
    `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '信箱',
    `phone` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '電話',
    `address` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '地址',
    `intro` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '教練介紹',
    `experience` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '年資',
    `motto` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '座右銘',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT `fk_trainers_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;