CREATE TABLE IF NOT EXISTS body_images (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `body_image`            VARCHAR(50) NOT NULL DEFAULT '' COMMENT '體態照片',
    `weight`                FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '體重(公斤)',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_body_images_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;