CREATE TABLE IF NOT EXISTS trainer_status_update_logs (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '紀錄 id',
    `user_id` INT(11) UNSIGNED COMMENT '教練 id',
    `trainer_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '教練狀態(1:正常/2:審核中/3:停權)',
    `comment` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '註解',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    CONSTRAINT fk_trainer_status_update_logs_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;