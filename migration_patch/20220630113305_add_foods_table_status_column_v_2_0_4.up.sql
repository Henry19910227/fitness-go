SET @dbname = "fitness";
SET @tablename = "foods";
SET @columnname = "status";
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
               " ADD ", @columnname, " TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '食物狀態(0:下架/1:上架)' AFTER amount_desc")
    )
);
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;