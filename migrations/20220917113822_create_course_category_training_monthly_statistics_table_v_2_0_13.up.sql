CREATE TABLE IF NOT EXISTS course_category_training_monthly_statistics (
    `id`        INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id',
    `category`  TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)',
    `year`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '年份',
    `month`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '月份',
    `total`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '當月總訓練數',
    `male`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '男性訓練數',
    `female`    INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '女性訓練數',
    `age_13_17` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '13-17歲訓練數',
    `age_18_24` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '18-24歲訓練數',
    `age_25_34` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '25-34歲訓練數',
    `age_35_44` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '35-44歲訓練數',
    `age_45_54` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '45-54歲訓練數',
    `age_55_64` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '55-64歲訓練數',
    `age_65_up` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '65+歲訓練數',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    UNIQUE KEY `unique_course_category_training_monthly_statistics` (`category`, `year`, `month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;