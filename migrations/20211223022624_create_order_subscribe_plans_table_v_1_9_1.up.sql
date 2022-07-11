CREATE TABLE IF NOT EXISTS order_subscribe_plans (
    `order_id`           VARCHAR(20) PRIMARY KEY COMMENT '訂單 id',
    `subscribe_plan_id`  INT(11) UNSIGNED COMMENT '訂閱方案 id',
    CONSTRAINT fk_order_subscribe_plans_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_order_subscribe_plans_subscribe_plan_id_to_subscribe_plans_id FOREIGN KEY (subscribe_plan_id) REFERENCES subscribe_plans(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;