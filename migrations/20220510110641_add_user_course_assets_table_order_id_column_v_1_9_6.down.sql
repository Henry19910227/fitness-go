SET @TableSchema = "fitness";
SET @TableName = "user_course_assets";
SET @Name = "order_id";
SET @DropColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_schema = @TableSchema)
      AND (table_name = @TableName)
      AND (column_name = @Name)
  ) > 0,
    CONCAT("ALTER TABLE", " ", @TableName, " ", "DROP", " ", @Name),
    "SELECT 1"
));
PREPARE DropColumn FROM @DropColumn;
EXECUTE DropColumn;
DEALLOCATE PREPARE DropColumn;