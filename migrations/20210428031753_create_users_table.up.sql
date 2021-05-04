CREATE TABLE IF NOT EXISTS users (
     `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
     `auth_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
     `account` VARCHAR(50) UNIQUE NOT NULL DEFAULT '',
     `password` VARCHAR(16) NOT NULL DEFAULT '',
     `device_token` VARCHAR(50) UNIQUE NOT NULL DEFAULT '',
     `user_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
     `user_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
     `create_at` DATETIME NOT NULL DEFAULT NOW(),
     `update_at` DATETIME NOT NULL DEFAULT NOW(),
     `email` VARCHAR(50) UNIQUE NOT NULL DEFAULT '',
     `nickname` VARCHAR(20) NOT NULL UNIQUE DEFAULT '',
     `sex` CHAR(1) NOT NULL DEFAULT 'm',
     `birthday` DATETIME NOT NULL DEFAULT NOW(),
     `height` tinyint(1) UNSIGNED NOT NULL DEFAULT '0',
     `weight` tinyint(1) UNSIGNED NOT NULL DEFAULT '0',
     `experience` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
     `target` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 10001;