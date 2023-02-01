CREATE TABLE IF NOT EXISTS workout_sets (
   `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '訓練組id',
   `workout_id` INT(11) UNSIGNED COMMENT '訓練id',
   `action_id` INT(11) UNSIGNED COMMENT '動作id',
   `type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '動作類別(1:動作/2:休息)',
   `auto_next` CHAR(1) NOT NULL DEFAULT 'N' COMMENT '自動下一組(Y:是/N:否)',
   `start_audio` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '前導語音',
   `progress_audio` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '進行中語音',
   `remark` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '備註',
   `weight` FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '重量(公斤)',
   `reps` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '次數',
   `distance` FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '距離(公里)',
   `duration` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '時長(秒)',
   `incline` FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '坡度',
   `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
   `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
   CONSTRAINT fk_workout_sets_workout_id_to_workouts_id FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
   CONSTRAINT fk_workout_sets_action_id_to_actions_id FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;