CREATE TABLE IF NOT EXISTS plans (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '計畫id',
    `course_id` INT(11) UNSIGNED NOT NULL COMMENT '課表id',
    `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '計畫名稱',
    `workout_count` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訓練數量',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT fk_plan_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB CHARSET=utf8mb4;