CREATE TABLE IF NOT EXISTS workout_set_orders (
    `workout_id` INT(11) UNSIGNED COMMENT '訓練id',
    `workout_set_id` INT(11) UNSIGNED COMMENT '訓練組id',
    `seq` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排列序號',
    PRIMARY KEY (`workout_id`,`workout_set_id`),
    CONSTRAINT fk_set_orders_workout_id_to_workouts_id FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
    CONSTRAINT fk_set_orders_workout_set_id_to_workout_sets_id FOREIGN KEY (workout_set_id) REFERENCES workout_sets(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4;