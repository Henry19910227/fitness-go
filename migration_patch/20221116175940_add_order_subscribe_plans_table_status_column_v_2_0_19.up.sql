SET @dbname = "fitness";
SET @tablename = "order_subscribe_plans";
SET @columnname = "status";
SET @preparedStatement = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = @columnname)
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD ", @columnname, " TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '訂閱狀態(0:無訂閱/1:訂閱中)' AFTER subscribe_plan_id")
    )
);
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;