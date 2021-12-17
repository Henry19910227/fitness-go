SET @dbname = "fitness";
SET @tablename = "courses";
SET @columnname = "sale_type";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = @columnname)
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD ", @columnname, " TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '銷售類型(1:免費課表/2:訂閱課表/3:付費課表)' AFTER user_id")
    )
);
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;