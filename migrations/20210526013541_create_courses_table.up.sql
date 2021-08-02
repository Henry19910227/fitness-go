CREATE TABLE IF NOT EXISTS courses (
   `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '課表 id',
   `user_id` INT(11) UNSIGNED NOT NULL COMMENT '用戶 id',
   `sale_id` INT(11) UNSIGNED COMMENT '銷售 id',
   `course_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)',
   `category` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)',
   `schedule_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排課類別(1:單一訓練/2:多項計畫)',
   `sale_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '銷售類型(1:免費課表/2:訂閱課表/3:付費課表)',
   `price` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '售價',
   `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '課表名稱',
   `cover` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '課表封面',
   `intro` VARCHAR(400) NOT NULL DEFAULT '' COMMENT '課表介紹',
   `food` VARCHAR(400) NOT NULL DEFAULT '' COMMENT '飲食建議',
   `level` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '強度(1:初級/2:中級/3:中高級/4:高級)',
   `suit` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)',
   `equipment` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)',
   `place` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)',
   `train_target` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)',
   `body_target` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)',
   `notice` VARCHAR(400) NOT NULL DEFAULT '' COMMENT '注意事項',
   `plan_count` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '計畫總數',
   `workout_count` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訓練總數',
   `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
   `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
   CONSTRAINT fk_courses_user_id_to_users_id FOREIGN KEY (user_id) REFERENCES users(id),
   CONSTRAINT fk_courses_sale_id_to_sale_items_id FOREIGN KEY (sale_id) REFERENCES sale_items(id)
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;