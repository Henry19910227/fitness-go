CREATE TABLE IF NOT EXISTS sale_items (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '銷售項目id',
    `type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '類型(1:免費課表/2:訂閱課表/3:付費課表)',
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '銷售名稱',
    `twd` INT(11) NOT NULL DEFAULT '0' COMMENT '台幣價格',
    `identifier` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '銷售識別碼',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間'
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;