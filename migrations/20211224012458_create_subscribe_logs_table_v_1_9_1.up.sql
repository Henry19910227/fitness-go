CREATE TABLE IF NOT EXISTS subscribe_logs (
    `id`           INT(11) UNSIGNED PRIMARY KEY COMMENT '紀錄 id',
    `user_id`      INT(11) UNSIGNED NOT NULL COMMENT '用戶 id',
    `order_id`     VARCHAR(20) NOT NULL COMMENT '訂單 id',
    `type`         TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '紀錄類型(1:訂閱/2:過期/3:退費)',
    `msg`          VARCHAR(50) NOT NULL DEFAULT '' COMMENT '紀錄訊息',
    `create_at`    DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_subscribe_logs_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_subscribe_logs_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci AUTO_INCREMENT = 1;