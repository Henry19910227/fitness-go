-- 請寫一個 Prepared Statement，
-- 檢查 user_course_statistics 是否存在 PRIMARY KEY `user_id` AND `course_id`，如果不存在就新增，
-- 除此之外，
-- 再寫另一個 Prepared Statement 檢查這個 PRIMARY KEY 是否存在，如果存在就刪除
SET @dbname = "fitness";
SET @tablename = "user_course_statistics";
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
        CONCAT("ALTER TABLE ", @tablename, " ADD PRIMARY KEY (`user_id`, `course_id`)")
    )
);
PREPARE addIfNotExists FROM @preparedStatement;
EXECUTE addIfNotExists;
DEALLOCATE PREPARE addIfNotExists;
