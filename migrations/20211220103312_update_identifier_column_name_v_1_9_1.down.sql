SET @dbname = "fitness";
SET @tablename = "sale_items";
SET @columnname = "product_id";
SET @preparedStatement = (
    SELECT IF (
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
            TABLE_NAME = @tablename AND TABLE_SCHEMA = @dbname AND COLUMN_NAME = @columnname
        ) > 0,
        CONCAT(" ALTER TABLE ", @tablename,
               " CHANGE COLUMN ", @columnname,
               " `identifier` varchar(100) NOT NULL DEFAULT '' COMMENT '銷售識別碼'"),
        "SELECT 1"
    )
);
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;