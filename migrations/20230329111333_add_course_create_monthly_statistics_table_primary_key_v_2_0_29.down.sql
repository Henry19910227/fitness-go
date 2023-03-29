SET @dbname = "fitness";
SET @tablename = "course_create_monthly_statistics";
SET @keyname = "PRIMARY";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (constraint_name = @keyname)
                AND (constraint_type = 'PRIMARY KEY')
        ) > 0,
        CONCAT("ALTER TABLE ", @tablename, " DROP PRIMARY KEY"),
        "SELECT 1"
    )
);
PREPARE dropIfExists FROM @preparedStatement;
EXECUTE dropIfExists;
DEALLOCATE PREPARE dropIfExists;