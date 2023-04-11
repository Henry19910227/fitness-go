SET @dbname = "fitness";
SET @tablename = "course_category_training_monthly_statistics";
SET @columnname = "id";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = @columnname)
        ) > 0,
        CONCAT("ALTER TABLE ", @tablename, " DROP COLUMN ", @columnname),
        "SELECT 1"
    )
);
PREPARE dropIfExists FROM @preparedStatement;
EXECUTE dropIfExists;
DEALLOCATE PREPARE dropIfExists;