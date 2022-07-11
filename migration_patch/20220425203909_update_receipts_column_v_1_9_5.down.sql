SET @dbname = "fitness";
SET @tablename = "receipts";

SET @updateOriginalTransactionIDColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "original_transaction_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " MODIFY COLUMN `original_transaction_id` varchar(20) NOT NULL DEFAULT '' COMMENT '初始交易 id'"),
    "SELECT 1"
));

SET @updateTransactionIDColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "transaction_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " MODIFY COLUMN `transaction_id` varchar(20) NOT NULL DEFAULT '' COMMENT '交易 id'"),
    "SELECT 1"
));

PREPARE updateOriginalTransactionIDColumn FROM @updateOriginalTransactionIDColumn;
EXECUTE updateOriginalTransactionIDColumn;
DEALLOCATE PREPARE updateOriginalTransactionIDColumn;

PREPARE updateTransactionIDColumn FROM @updateTransactionIDColumn;
EXECUTE updateTransactionIDColumn;
DEALLOCATE PREPARE updateTransactionIDColumn;