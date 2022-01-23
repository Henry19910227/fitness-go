CREATE TABLE IF NOT EXISTS subscribe_logs (
    `id`             INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '訂閱紀錄id',
    `user_id`        INT(11) UNSIGNED COMMENT '用戶 id',
    `transaction_id` VARCHAR(20) COMMENT '交易 id',
    `purchase_date`  DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱購買日期',
    `expires_date`   DATETIME NOT NULL DEFAULT NOW() COMMENT '訂閱過期日期',
    `type`           VARCHAR(20) NOT NULL DEFAULT '' COMMENT '紀錄類型(初次訂閱:initial_buy/恢復訂閱:resubscribe/續訂:renew/訂閱升級:upgrade/訂閱降級:downgrade/訂閱過期:expired/退費:refund)',
    `msg`            VARCHAR(50) NOT NULL DEFAULT '' COMMENT '紀錄訊息',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_subscribe_logs_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_subscribe_logs_transaction_id_to_receipts_transaction_id FOREIGN KEY (transaction_id) REFERENCES receipts(transaction_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci AUTO_INCREMENT = 1;