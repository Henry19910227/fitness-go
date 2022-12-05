SET @dbname = "fitness";
SET @tablename = "user_course_assets";
SET @columnname = "source";
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
               " ADD ", @columnname, " TINYINT(1) NOT NULL DEFAULT '1' COMMENT '來源(0:未知/1:購買/2:贈送)' AFTER available")
    )
);
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;