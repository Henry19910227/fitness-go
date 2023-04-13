SET @dbname = "fitness";
SET @tablename = "course_training_avg_statistics";
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
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename, " ADD PRIMARY KEY (`course_id`)")
    )
);
PREPARE addIfNotExists FROM @preparedStatement;
EXECUTE addIfNotExists;
DEALLOCATE PREPARE addIfNotExists;