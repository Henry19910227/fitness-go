CREATE TABLE IF NOT EXISTS purchase_logs (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '購買紀錄id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶 id',
    `order_id`       VARCHAR(20) COMMENT '訂單 id',
    `type`           TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '訂單類型(1:購買/2:退費)',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_purchase_logs_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_purchase_logs_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;