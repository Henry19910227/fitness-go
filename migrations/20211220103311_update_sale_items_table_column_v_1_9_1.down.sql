SET @dbname = "fitness";
SET @tablename = "sale_items";

SET @addTwdColumn = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = "twd")
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD COLUMN `twd` INT(11) NOT NULL DEFAULT '0' COMMENT '台幣價格' AFTER `type`")
    )
);

SET @addIdentifierColumn = (SELECT IF(
        (
            SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
            WHERE
                (table_name = @tablename)
                AND (table_schema = @dbname)
                AND (column_name = "identifier")
        ) > 0,
        "SELECT 1",
        CONCAT("ALTER TABLE ", @tablename,
               " ADD COLUMN `identifier` varchar(100) NOT NULL DEFAULT '' COMMENT '產品id' AFTER `type`")
    )
);

SET @dropEnableColumn = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "enable")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP COLUMN `enable`"),
    "SELECT 1"
));

SET @dropProductLabelIdForeignKey = (SELECT IF(
  (
    SELECT COUNT(*) FROM information_schema.REFERENTIAL_CONSTRAINTS
    WHERE
      (TABLE_NAME = @tablename)
      AND (CONSTRAINT_SCHEMA = @dbname)
      AND (CONSTRAINT_NAME = "fk_sale_items_label_id_to_product_labels_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP FOREIGN KEY `fk_sale_items_label_id_to_product_labels_id`"),
    "SELECT 1"
));

SET @dropProductLabelIdIndex = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS
    WHERE
      (TABLE_NAME = @tablename)
      AND (table_schema = @dbname)
      AND (INDEX_NAME = "fk_sale_items_label_id_to_product_labels_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP INDEX `fk_sale_items_label_id_to_product_labels_id`"),
    "SELECT 1"
));

SET @dropProductLabelId = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "product_label_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP COLUMN `product_label_id`"),
    "SELECT 1"
));

PREPARE addTwdColumn FROM @addTwdColumn;
EXECUTE addTwdColumn;
DEALLOCATE PREPARE addTwdColumn;

PREPARE addIdentifierColumn FROM @addIdentifierColumn;
EXECUTE addIdentifierColumn;
DEALLOCATE PREPARE addIdentifierColumn;

PREPARE dropEnableColumn FROM @dropEnableColumn;
EXECUTE dropEnableColumn;
DEALLOCATE PREPARE dropEnableColumn;

PREPARE dropProductLabelIdForeignKey FROM @dropProductLabelIdForeignKey;
EXECUTE dropProductLabelIdForeignKey;
DEALLOCATE PREPARE dropProductLabelIdForeignKey;

PREPARE dropProductLabelIdIndex FROM @dropProductLabelIdIndex;
EXECUTE dropProductLabelIdIndex;
DEALLOCATE PREPARE dropProductLabelIdIndex;

PREPARE dropProductLabelId FROM @dropProductLabelId;
EXECUTE dropProductLabelId;
DEALLOCATE PREPARE dropProductLabelId;