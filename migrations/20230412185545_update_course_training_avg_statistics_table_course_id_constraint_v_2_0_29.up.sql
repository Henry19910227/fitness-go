SET @dbname = "fitness";
SET @tablename = "course_training_avg_statistics";
SET @key = "fk_course_training_avg_statistics_course_id_to_courses_id";

SET @dropForeignKey = (SELECT IF(
  (
    SELECT COUNT(*) FROM information_schema.REFERENTIAL_CONSTRAINTS
    WHERE
      (TABLE_NAME = @tablename)
      AND (CONSTRAINT_SCHEMA = @dbname)
      AND (CONSTRAINT_NAME = @key)
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP FOREIGN KEY " ,@key),
    "SELECT 1"
));


SET @addForeignKey = CONCAT("ALTER TABLE ", @tablename, " ADD CONSTRAINT ", @key, " FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE");

PREPARE dropForeignKey FROM @dropForeignKey;
EXECUTE dropForeignKey;
DEALLOCATE PREPARE dropForeignKey;

PREPARE addForeignKey FROM @addForeignKey;
EXECUTE addForeignKey;
DEALLOCATE PREPARE addForeignKey;