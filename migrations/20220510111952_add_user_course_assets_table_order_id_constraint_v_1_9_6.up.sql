SET @TableSchema = "fitness";
SET @TableName = "user_course_assets";
SET @Type = "CONSTRAINT";
SET @Name = "fk_purchase_courses_order_id_to_orders_id";
SET @Attribute = "FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE";
SET @AddConstraint = (SELECT IF(
      (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.`referential_constraints`
        WHERE
          (CONSTRAINT_SCHEMA = @TableSchema)
          AND (TABLE_NAME = @TableName)
          AND (CONSTRAINT_NAME = @Name)
      ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE", " ", @TableName, " ",
               "ADD", " ", @Type, " ", @Name, " ", @Attribute)
    )
);
PREPARE AddConstraint FROM @AddConstraint;
EXECUTE AddConstraint;
DEALLOCATE PREPARE AddConstraint;