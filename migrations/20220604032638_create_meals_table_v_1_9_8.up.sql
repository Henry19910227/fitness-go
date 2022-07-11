CREATE TABLE IF NOT EXISTS meals (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `diet_id`               INT(11) UNSIGNED COMMENT '飲食紀錄id',
    `food_id`               INT(11) UNSIGNED COMMENT '食物id',
    `type`                  TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '類型(1:/早餐/2:午餐/3:晚餐/4:點心)',
    `amount`                FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '數量',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    UNIQUE KEY `unique_meals` (`diet_id`, `food_id`, `type`),
    CONSTRAINT fk_meals_diet_id_to_diets_id FOREIGN KEY (diet_id) REFERENCES diets(id) ON DELETE CASCADE,
    CONSTRAINT fk_meals_food_id_to_foods_id FOREIGN KEY (food_id) REFERENCES foods(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;