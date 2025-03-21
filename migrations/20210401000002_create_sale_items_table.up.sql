CREATE TABLE IF NOT EXISTS sale_items (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '銷售項目id',
    `product_label_id` INT(11) UNSIGNED COMMENT '產品標籤id',
    `type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '類型(1:免費課表/3:付費課表)',
    `enable` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否啟用',
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '銷售名稱',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT `fk_sale_items_product_label_id_to_product_labels_id` FOREIGN KEY (`product_label_id`) REFERENCES `product_labels` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;