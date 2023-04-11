SET @dbname = "fitness";
SET @tablename = "course_category_training_monthly_statistics";
SET @keyname = "unique_course_category_training_monthly_statistics";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (index_name = @keyname)
                AND (non_unique = 0)
        ) > 0,
        CONCAT("ALTER TABLE ", @tablename, " DROP INDEX ", @keyname),
        "SELECT 1"
    )
);
PREPARE dropIfExists FROM @preparedStatement;
EXECUTE dropIfExists;
DEALLOCATE PREPARE dropIfExists;