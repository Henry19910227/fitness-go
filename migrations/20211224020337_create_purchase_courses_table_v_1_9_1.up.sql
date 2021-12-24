CREATE TABLE IF NOT EXISTS purchase_courses (
    `id`           INT(11) UNSIGNED PRIMARY KEY COMMENT '紀錄 id',
    `user_id`      INT(11) UNSIGNED NOT NULL COMMENT '用戶 id',
    `order_id`     VARCHAR(20) NOT NULL COMMENT '訂單 id',
    `course_id`    INT(11) UNSIGNED NOT NULL COMMENT '課表 id',
    `create_at`    DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_purchase_courses_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_purchase_courses_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT fk_purchase_courses_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci AUTO_INCREMENT = 1;