CREATE TABLE IF NOT EXISTS favorite_trainers (
    `user_id`        INT(11) UNSIGNED COMMENT '用戶id',
    `trainer_id`      INT(11) UNSIGNED COMMENT '教練id',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    PRIMARY KEY (`user_id`,`trainer_id`),
    CONSTRAINT `fk_favorite_trainers_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_favorite_trainers_trainer_id_to_users_id` FOREIGN KEY (`trainer_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;