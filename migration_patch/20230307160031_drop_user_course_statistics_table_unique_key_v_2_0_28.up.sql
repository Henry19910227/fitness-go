-- 請寫一個 Prepared Statement，
-- 檢查 user_course_statistics 是否存在 UNIQUE KEY `unique_user_course_statistics`，如果存在就刪除，
-- 除此之外，再寫另一個 Prepared Statement 檢查 user_course_statistics 是否存在 UNIQUE KEY `unique_user_course_statistics`，
-- 如果不存在就加回來
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
        CONCAT("ALTER TABLE ", @tablename, " DROP INDEX ", @keyname),
        "SELECT 1"
    )
);
PREPARE dropIfExists FROM @preparedStatement;
EXECUTE dropIfExists;
DEALLOCATE PREPARE dropIfExists;
