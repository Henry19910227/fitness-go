SET @TableSchema = "fitness";
SET @TableName = "user_course_assets";
SET @Name = "fk_purchase_courses_order_id_to_orders_id";
SET @DropConstraint = (SELECT IF(
      (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.`KEY_COLUMN_USAGE`
        WHERE
          (TABLE_SCHEMA = @TableSchema)
          AND (TABLE_NAME = @TableName)
          AND (CONSTRAINT_NAME = @Name)
      ) > 0,
        CONCAT("ALTER TABLE", " ", @TableName, " ", "DROP FOREIGN KEY", " ", @Name),
        "SELECT 1"
));
SET @DropIndex = (SELECT IF(
      (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.`STATISTICS`
        WHERE
          (TABLE_SCHEMA = @TableSchema)
          AND (TABLE_NAME = @TableName)
          AND (INDEX_NAME = @Name)
      ) > 0,
        CONCAT("ALTER TABLE", " ", @TableName, " ", "DROP INDEX", " ", @Name),
        "SELECT 1"
));
PREPARE DropConstraint FROM @DropConstraint;
EXECUTE DropConstraint;
DEALLOCATE PREPARE DropConstraint;

PREPARE DropIndex FROM @DropIndex;
EXECUTE DropIndex;
DEALLOCATE PREPARE DropIndex;