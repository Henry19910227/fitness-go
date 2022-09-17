CREATE TABLE IF NOT EXISTS course_training_monthly_statistics (
    `id`        INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id',
    `year`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '年份',
    `month`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '月份',
    `total`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '當月總訓練數',
    `free`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '免費課表訓練數',
    `subscribe`    INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訂閱課表訓練數',
    `charge`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '付費課表訓練數',
    `aerobic`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '有氧課表訓練數',
    `interval_training`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '間歇肌力訓練課表訓練數',
    `weight_training`    INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '重量訓練課表訓練數',
    `resistance_training` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '阻力訓練課表訓練數',
    `bodyweight_training` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '徒手訓練課表訓練數',
    `other_training` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '其他訓練課表訓練數',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_course_training_monthly_statistics` (`year`, `month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;