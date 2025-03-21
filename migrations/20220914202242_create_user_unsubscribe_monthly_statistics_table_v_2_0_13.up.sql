CREATE TABLE IF NOT EXISTS user_unsubscribe_monthly_statistics (
    `year`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '年份',
    `month`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '月份',
    `total`     INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '當月總退訂人數',
    `male`      INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '男性退訂人數',
    `female`    INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '女性退訂人數',
    `age_13_17` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '13-17歲退訂人數',
    `age_18_24` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '18-24歲退訂人數',
    `age_25_34` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '25-34歲退訂人數',
    `age_35_44` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '35-44歲退訂人數',
    `age_45_54` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '45-54歲退訂人數',
    `age_55_64` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '55-64歲退訂人數',
    `age_65_up` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '65+歲退訂人數',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新日期',
    PRIMARY KEY (`year`,`month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT = 1;