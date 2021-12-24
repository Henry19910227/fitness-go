CREATE TABLE IF NOT EXISTS members (
    `user_id` INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `member_status` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '會員狀態(1:正常/2:到期/3:取消)',
    `start_date` DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱開始日期',
    `end_date` DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱過期日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT `fk_members_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;