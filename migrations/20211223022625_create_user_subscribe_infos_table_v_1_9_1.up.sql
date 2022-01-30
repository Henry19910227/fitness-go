CREATE TABLE IF NOT EXISTS user_subscribe_infos (
    `user_id`           INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `subscribe_plan_id` INT(11) UNSIGNED COMMENT '訂閱方案id',
    `status`            TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '會員狀態(0:無會員/1:付費會員)',
    `start_date`        DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱開始日期',
    `expires_date`      DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱過期日期',
    `update_at`         DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT `fk_subscribe_infos_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_subscribe_infos_plan_id_to_subscribe_plans_id` FOREIGN KEY (`subscribe_plan_id`) REFERENCES `subscribe_plans` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;