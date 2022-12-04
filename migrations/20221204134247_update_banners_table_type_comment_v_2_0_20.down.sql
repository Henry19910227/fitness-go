SET @dbname = "fitness";
SET @tablename = "banners";

SET @preparedStatement = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "type")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " MODIFY COLUMN `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '類型(1:課表/2:教練/3:訂閱)'"),
    "SELECT 1"
));

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;