CREATE TABLE IF NOT EXISTS foods (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `food_category_id`      INT(11) UNSIGNED COMMENT '食物類別id',
    `source`                TINYINT(1) NOT NULL DEFAULT '1' COMMENT '動作來源(1:系統創建食物/2:用戶創建食物',
    `name`                  VARCHAR(50) NOT NULL DEFAULT '' COMMENT '食物名稱',
    `calorie`               INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '食物熱量',
    `amount_desc`           VARCHAR(50) NOT NULL DEFAULT '' COMMENT '份量描述',
    `status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '動作狀態(0:下架/1:上架)',
    `is_deleted`             TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否刪除',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_foods_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_foods_food_category_id_to_food_categories_id FOREIGN KEY (food_category_id) REFERENCES food_categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;