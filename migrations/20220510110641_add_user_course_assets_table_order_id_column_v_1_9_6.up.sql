SET @TableSchema = "fitness";
SET @TableName = "user_course_assets";
SET @Type = "COLUMN";
SET @Name = "order_id";
SET @Attribute = "VARCHAR(20) COMMENT '訂單 id' AFTER id";
SET @AddColumn = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_schema = @TableSchema)
                AND (table_name = @TableName)
                AND (column_name = @Name)
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE", " ", @TableName, " ",
               "ADD", " ", @Type, " ", @Name, " ", @Attribute)
    )
);
PREPARE AddColumn FROM @AddColumn;
EXECUTE AddColumn;
DEALLOCATE PREPARE AddColumn;