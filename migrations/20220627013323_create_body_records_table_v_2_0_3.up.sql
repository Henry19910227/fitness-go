CREATE TABLE IF NOT EXISTS body_records (
    `id`                    INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主鍵id',
    `user_id`               INT(11) UNSIGNED COMMENT '用戶id',
    `record_type`           TINYINT(2) UNSIGNED NOT NULL DEFAULT '1' COMMENT '紀錄類型(1:體重紀錄/2:體脂紀錄/3:胸圍紀錄/4:腰圍紀錄/5:臀圍紀錄/6:身高紀錄/7:臂圍紀錄/8:小臂圍紀錄/9:肩圍紀錄/10:大腿圍紀錄/11:小腿圍紀錄/12:頸圍紀錄)',
    `value`                 FLOAT UNSIGNED NOT NULL DEFAULT '0' COMMENT '數值',
    `create_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at`             DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    CONSTRAINT fk_body_records_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;