SET @dbname = "fitness";
SET @tablename = "sale_items";

SET @deleteIdentifierColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "identifier")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP COLUMN `identifier`"),
    "SELECT 1"
));

SET @addProductLabelIdColumn = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = "product_label_id")
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD COLUMN `product_label_id` INT(11) UNSIGNED COMMENT '產品標籤id' AFTER id")
    )
);

SET @addProductLabelIdConstraint = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = "fk_sale_items_label_id_to_product_labels_id")
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD CONSTRAINT `fk_sale_items_label_id_to_product_labels_id` FOREIGN KEY (product_label_id) REFERENCES product_labels(id) ON DELETE CASCADE")
    )
);

SET @addEnableColumn = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = "enable")
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD COLUMN `enable` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否啟用' AFTER `type`")
    )
);

SET @deleteTwdColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "twd")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP COLUMN `twd`"),
    "SELECT 1"
));

PREPARE deleteIdentifierColumn FROM @deleteIdentifierColumn;
EXECUTE deleteIdentifierColumn;
DEALLOCATE PREPARE deleteIdentifierColumn;

PREPARE addProductLabelIdColumn FROM @addProductLabelIdColumn;
EXECUTE addProductLabelIdColumn;
DEALLOCATE PREPARE addProductLabelIdColumn;

PREPARE addProductLabelIdConstraint FROM @addProductLabelIdConstraint;
EXECUTE addProductLabelIdConstraint;
DEALLOCATE PREPARE addProductLabelIdConstraint;

PREPARE addEnableColumn FROM @addEnableColumn;
EXECUTE addEnableColumn;
DEALLOCATE PREPARE addEnableColumn;

PREPARE deleteTwdColumn FROM @deleteTwdColumn;
EXECUTE deleteTwdColumn;
DEALLOCATE PREPARE deleteTwdColumn;
















