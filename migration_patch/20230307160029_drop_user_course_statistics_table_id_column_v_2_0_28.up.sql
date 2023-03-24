-- 請寫一個 Prepared Statement，
-- 檢查 user_course_statistics 是否存在 id 欄位，如果存在就刪除，
-- 除此之外，
-- 再寫另一個 Prepared Statement 檢查 user_course_statistics 是否存在 id 欄位，
-- 如果不存在就加回來且該欄位為 primary key
SET @dbname = "fitness";
SET @tablename = "user_course_statistics";
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