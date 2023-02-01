SET @dbname = "fitness";
SET @tablename = "workout_logs";

SET @updateWorkoutLogsPlaceComment = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (table_name = @tablename)
      AND (table_schema = @dbname)
      AND (column_name = "place")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " MODIFY COLUMN `place` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '地點(0:未指定/1:住家/2:健身房/3:戶外)'"),
    "SELECT 1"
));

PREPARE updateWorkoutLogsPlaceComment FROM @updateWorkoutLogsPlaceComment;
EXECUTE updateWorkoutLogsPlaceComment;
DEALLOCATE PREPARE updateWorkoutLogsPlaceComment;