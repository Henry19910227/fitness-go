CREATE TABLE IF NOT EXISTS actions (
    `id` INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '動作id',
    `course_id` INT(11) UNSIGNED COMMENT '課表id',
    `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '動作名稱',
    `source` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '動作來源(1:系統動作/2:教練自創動作)',
    `type` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)',
    `category` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)',
    `body` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)',
    `equipment` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)',
    `intro` VARCHAR(400) NOT NULL DEFAULT '' COMMENT '動作介紹',
    `cover` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '封面',
    `video` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '動作影片',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建時間',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新時間',
    CONSTRAINT fk_actions_course_id_to_courses_id FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB CHARSET=utf8mb4 AUTO_INCREMENT = 1;