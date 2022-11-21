SET @dbname = "fitness";
SET @tablename = "user_subscribe_infos";
SET @columnname = "original_transaction_id";
SET @preparedStatement = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = @columnname)
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP ", @columnname, ";"),
    "SELECT 1"
));
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;