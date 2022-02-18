CREATE TABLE IF NOT EXISTS order_courses (
    `order_id`     VARCHAR(20) PRIMARY KEY COMMENT '訂單 id',
    `sale_item_id` INT(11) UNSIGNED COMMENT '銷售項目 id',
    `course_id`    INT(11) UNSIGNED COMMENT '課表 id',
    CONSTRAINT fk_order_courses_order_id_to_orders_id FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT fk_order_courses_item_id_to_sale_items_id FOREIGN KEY (sale_item_id) REFERENCES sale_items(id),
    CONSTRAINT fk_order_courses_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;