CREATE TABLE IF NOT EXISTS favorite_actions (
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `action_id`      INT(11) UNSIGNED COMMENT '動作id',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    PRIMARY KEY (`user_id`,`action_id`),
    CONSTRAINT `fk_favorite_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_favorite_action_id_to_actions_id` FOREIGN KEY (`action_id`) REFERENCES `actions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;