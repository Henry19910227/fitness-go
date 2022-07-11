CREATE TABLE IF NOT EXISTS subscribe_plans (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '訂閱方案id',
    `product_label_id` INT(11) UNSIGNED COMMENT '產品標籤id',
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '訂閱名稱',
    `period` TINYINT(1) UNSIGNED NOT NULL DEFAULT '99' COMMENT '週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)',
    `enable` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否啟用',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT fk_subscribe_plans_label_id_to_product_labels_id FOREIGN KEY (product_label_id) REFERENCES product_labels(id) ON DELETE SET NULL
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;