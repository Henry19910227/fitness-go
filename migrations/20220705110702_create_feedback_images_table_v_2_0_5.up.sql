CREATE TABLE IF NOT EXISTS feedback_images (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `feedback_id`           INT(11) UNSIGNED COMMENT '反饋id',
    `image`                 VARCHAR(50) NOT NULL DEFAULT '' COMMENT '反饋照片',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_feedback_images_feedback_id_to_feedbacks_id FOREIGN KEY (feedback_id) REFERENCES feedbacks(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;