CREATE TABLE IF NOT EXISTS receipts (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '收據id',
    `order_id`     VARCHAR(20) DEFAULT '' COMMENT '訂單 id',
    `payment_type` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '支付方式(1:apple內購/2:google內購)',
    `receipt_token` TEXT NOT NULL COMMENT '收據 token',
    `original_transaction_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '初始交易 id',
    `transaction_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '交易 id',
    `product_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '銷售 id',
    `quantity` INT(11) NOT NULL DEFAULT '1' COMMENT '數量',
    `create_at`    DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    INDEX(transaction_id),
    UNIQUE KEY `unique_receipts` (`order_id`,`original_transaction_id`,`transaction_id`),
    CONSTRAINT fk_receipts_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;