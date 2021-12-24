CREATE TABLE IF NOT EXISTS orders (
    `id`           VARCHAR(20) PRIMARY KEY COMMENT '訂單 id',
    `user_id`      INT(11) UNSIGNED NOT NULL COMMENT '用戶 id',
    `sale_item_id` INT(11) UNSIGNED NOT NULL COMMENT '銷售項目 id',
    `quantity`     INT(11) UNSIGNED NOT NULL DEFAULT '1' COMMENT '數量',
    `order_type`   TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '訂單類型(1:課表購買/2:會員訂閱)',
    `payment_type` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '支付方式(1:apple內購/2:google內購)',
    `order_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)',
    `create_at`    DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at`    DATETIME NOT NULL DEFAULT NOW() COMMENT '修改時間',
    CONSTRAINT fk_orders_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_orders_sale_item_id_to_sale_items_id FOREIGN KEY (sale_item_id) REFERENCES sale_items(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;