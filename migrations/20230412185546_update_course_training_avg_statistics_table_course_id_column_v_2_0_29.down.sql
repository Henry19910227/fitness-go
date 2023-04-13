SET @dbname = "fitness";
SET @tablename = "course_training_avg_statistics";

SET @preparedStatement = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "course_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " MODIFY COLUMN `course_id` INT(11) UNSIGNED COMMENT '課表id'"),
    "SELECT 1"
));

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;