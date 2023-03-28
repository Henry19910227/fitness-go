SET @dbname = "fitness";
SET @tablename = "user_unsubscribe_monthly_statistics";
SET @columnname = "id";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = @columnname)
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename, " ADD COLUMN ", @columnname, " INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '報表id' FIRST")
    )
);
PREPARE addIfNotExists FROM @preparedStatement;
EXECUTE addIfNotExists;
DEALLOCATE PREPARE addIfNotExists;