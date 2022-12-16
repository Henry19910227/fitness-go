CREATE TABLE IF NOT EXISTS ios_versions (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'ios 版本號',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間'
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;