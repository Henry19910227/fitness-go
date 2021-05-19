CREATE TABLE IF NOT EXISTS users (
     `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '用戶id',
     `account_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)',
     `account` VARCHAR(50) UNIQUE NOT NULL DEFAULT '' COMMENT '帳號',
     `password` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '密碼',
     `device_token` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '推播token',
     `user_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用戶狀態 (1:正常/2:違規/3:刪除)',
     `user_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用戶類型 (1:一般用戶/2:訂閱用戶)',
     `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
     `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '修改日期',
     `email` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '信箱',
     `nickname` VARCHAR(20) NOT NULL UNIQUE DEFAULT '' COMMENT '暱稱',
     `sex` CHAR(1) NOT NULL DEFAULT '' COMMENT '性別 (m:男/f:女)',
     `birthday` DATE NOT NULL DEFAULT '0000-01-01' COMMENT '生日',
     `height` FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '身高',
     `weight` FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '體重',
     `experience` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)',
     `target` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '目標 (0:未指定/1:減重/2:維持健康/3:增肌)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 10001;