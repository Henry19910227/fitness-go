SET @dbname = "fitness";
SET @tablename = "users";

SET @update_column = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "device_token")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " MODIFY COLUMN `device_token` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '推播token'"),
    "SELECT 1"
));

PREPARE update_column FROM @update_column;
EXECUTE update_column;
DEALLOCATE PREPARE update_column;