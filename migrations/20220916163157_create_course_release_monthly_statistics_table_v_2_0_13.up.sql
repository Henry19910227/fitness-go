CREATE TABLE IF NOT EXISTS course_release_monthly_statistics (
    `id`        INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id',
    `year`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '年份',
    `month`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '月份',
    `total`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '當月總上架數',
    `free`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '免費課表上架數',
    `subscribe`    INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訂閱課表上架數',
    `charge`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '付費課表上架數',
    `aerobic`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '有氧課表上架數',
    `interval_training`  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '間歇肌力上架課表上架數',
    `weight_training`    INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '重量上架課表上架數',
    `resistance_training` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '阻力上架課表上架數',
    `bodyweight_training` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '徒手上架課表上架數',
    `other_training` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '其他上架課表上架數',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_course_release_monthly_statistics` (`year`, `month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;