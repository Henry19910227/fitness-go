CREATE TABLE IF NOT EXISTS bank_accounts (
    `user_id` INT(11) UNSIGNED PRIMARY KEY COMMENT '用戶id',
    `account_name` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '帳戶名',
    `account` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '帳戶',
    `account_image` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '帳戶照片',
    `bank_code` VARCHAR(3) NOT NULL DEFAULT '' COMMENT '銀行代號',
    `branch` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '分行',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_bank_accounts_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;