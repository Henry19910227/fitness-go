CREATE TABLE IF NOT EXISTS user_subscribe_infos (
    `user_id`           INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `order_id`          VARCHAR(20) COMMENT '訂單 id',
    `original_transaction_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '當前綁定的交易 id',
    `status`            TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '會員狀態(0:無會員/1:付費會員)',
    `payment_type`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '當前支付方式(0:未指定/1:apple內購/2:google內購)',
    `start_date`        DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱開始日期',
    `expires_date`      DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱過期日期',
    `update_at`         DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT `fk_subscribe_infos_user_id_to_users_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_subscribe_infos_order_id_to_orders_id` FOREIGN KEY (`order_id`) REFERENCES `orders`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;