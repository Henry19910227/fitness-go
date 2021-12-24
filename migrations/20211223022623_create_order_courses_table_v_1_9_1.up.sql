CREATE TABLE IF NOT EXISTS order_courses (
    `order_id`     VARCHAR(20) PRIMARY KEY COMMENT '訂單 id',
    `course_id`    INT(11) UNSIGNED NOT NULL COMMENT '課表 id',
    CONSTRAINT fk_order_courses_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT fk_order_courses_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;