CREATE TABLE IF NOT EXISTS review_images (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '圖片id',
    `review_id` INT(11) UNSIGNED COMMENT '評論id',
    `image` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '圖片',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    CONSTRAINT fk_review_images_review_id_to_reviews_id FOREIGN KEY (review_id) REFERENCES reviews(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;