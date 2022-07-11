SET @dbname = "fitness";
SET @tablename = "courses";
SET @columnname = "sale_type";
SET @preparedStatement = (SELECT IF(
			(
				SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
				WHERE
					(table_name = @tablename)
					AND (table_schema = @dbname)
					AND (column_name = @columnname)
			) > 0,
            CONCAT("UPDATE ", @tablename,
                   " INNER JOIN sale_items ON courses.sale_id = sale_items.id",
                   " SET courses.sale_type = sale_items.type"),
			'SELECT 1'
		)
);
PREPARE fillData FROM @preparedStatement;
EXECUTE fillData;
DEALLOCATE PREPARE fillData;