SET @dbname = "fitness";
SET @tablename = "food_categories";

SET @deleteOriginalTransactionIDColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
  ) > 0,
    CONCAT(" DELETE FROM ", @tablename),
    "SELECT 1"
));

PREPARE deleteOriginalTransactionIDColumn FROM @deleteOriginalTransactionIDColumn;
EXECUTE deleteOriginalTransactionIDColumn;
DEALLOCATE PREPARE deleteOriginalTransactionIDColumn;