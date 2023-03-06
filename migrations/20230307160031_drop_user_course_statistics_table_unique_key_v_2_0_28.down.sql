SET @dbname = "fitness";
SET @tablename = "user_course_statistics";
SET @keyname = "unique_user_course_statistics";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (index_name = @keyname)
                AND (non_unique = 0)
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename, " ADD UNIQUE KEY ", @keyname, " (`user_id`,`course_id`)")
    )
);
PREPARE addIfNotExists FROM @preparedStatement;
EXECUTE addIfNotExists;
DEALLOCATE PREPARE addIfNotExists;
