SET @dbname = "fitness";
SET @tablename = "sale_items";
SET @columnname = "twd";
SET @preparedStatement = (
    SELECT IF(
        (
            SELECT DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
            TABLE_NAME = @tablename AND TABLE_SCHEMA = @dbname AND COLUMN_NAME = @columnname
        ) = 'int',
        CONCAT(" ALTER TABLE ", @tablename,
               " MODIFY COLUMN ", @columnname,
               " DECIMAL(7,2) NOT NULL DEFAULT '0' COMMENT '台幣價格'"),
        "SELECT 1"
    )
);
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;