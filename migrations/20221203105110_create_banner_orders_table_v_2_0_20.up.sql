CREATE TABLE IF NOT EXISTS banner_orders (
    `banner_id` INT(11) UNSIGNED COMMENT '訓練id',
    `seq` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排列序號',
    PRIMARY KEY (`banner_id`),
    CONSTRAINT fk_banner_orders_banner_id_to_banners_id FOREIGN KEY (banner_id) REFERENCES banners(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4;