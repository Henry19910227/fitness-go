SET @dbname = "fitness";
SET @tablename = "food_categories";

SET @insertFoodCategoriesValues = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
  ) > 0,
    CONCAT(" INSERT INTO ", @tablename, " (id,tag,title,update_at) VALUES ",
    " (1,1,'米麥類',NOW()), ",
    " (2,1,'根莖類',NOW()), ",
    " (3,1,'其他',NOW()), ",

    " (4,2,'豆類',NOW()), ",
    " (5,2,'家畜',NOW()), ",
    " (6,2,'家禽',NOW()), ",
    " (7,2,'蛋類',NOW()), ",
    " (8,2,'水產',NOW()), ",
    " (9,2,'其他',NOW()), ",

    " (10,3,'柑橘類',NOW()), ",
    " (11,3,'蘋果類',NOW()), ",
    " (12,3,'瓜類',NOW()), ",
    " (13,3,'芒果類',NOW()), ",
    " (14,3,'芭樂類',NOW()), ",
    " (15,3,'梨類',NOW()), ",
    " (16,3,'桃類',NOW()), ",
    " (17,3,'柿類',NOW()), ",
    " (18,3,'李類',NOW()), ",
    " (19,3,'棗類',NOW()), ",
    " (20,3,'果汁類',NOW()), ",
    " (21,3,'其他',NOW()), ",

    " (22,4,'一般蔬菜',NOW()), ",
    " (23,4,'高蛋白蔬菜',NOW()), ",
    " (24,4,'其他',NOW()), ",

    " (25,5,'低脂奶類',NOW()), ",
    " (26,5,'脫脂奶類',NOW()), ",
    " (27,5,'其他',NOW()), ",

    " (28,6,'植物類',NOW()), ",
    " (29,6,'堅果類',NOW()), ",
    " (30,6,'其他',NOW()) ",

    " ON DUPLICATE KEY UPDATE "
    " tag=VALUES(tag), ",
    " title=VALUES(title), ",
    " update_at=VALUES(update_at) "),
    "SELECT 1"
));

PREPARE insertFoodCategoriesValues FROM @insertFoodCategoriesValues;
EXECUTE insertFoodCategoriesValues;
DEALLOCATE PREPARE insertFoodCategoriesValues;