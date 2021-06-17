CREATE TABLE IF NOT EXISTS workouts (
     `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '訓練id',
     `plan_id` INT(11) UNSIGNED NOT NULL COMMENT '計畫id',
     `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '訓練名稱',
     `equipment` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)',
     `start_audio` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '前導語音',
     `end_audio` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '結束語音',
     `workout_set_count` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '動作組數',
     `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
     `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
     CONSTRAINT fk_workouts_plan_id_to_plans_id FOREIGN KEY (plan_id) REFERENCES plans(id)
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;