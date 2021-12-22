CREATE TABLE IF NOT EXISTS receipts (
    `id`           VARCHAR(20) PRIMARY KEY COMMENT '收據 id',
    `order_id`     VARCHAR(20) NOT NULL DEFAULT '' COMMENT '訂單 id',
    `receipt_token` TEXT NOT NULL COMMENT '收據 token',
    `original_transaction_id` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '初始交易 id',
    `transaction_id` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '交易 id',
    `product_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '銷售 id',
    `quantity` INT(11) NOT NULL DEFAULT '1' COMMENT '數量',
    `order_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)',
    `original_purchase_date` DATETIME NOT NULL DEFAULT NOW() COMMENT '首次支付日期',
    `purchase_date` DATETIME NOT NULL DEFAULT NOW() COMMENT '支付日期',
    `expires_date` DATETIME NOT NULL DEFAULT NOW() COMMENT '到期日期',
    CONSTRAINT fk_receipts_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci AUTO_INCREMENT = 1;