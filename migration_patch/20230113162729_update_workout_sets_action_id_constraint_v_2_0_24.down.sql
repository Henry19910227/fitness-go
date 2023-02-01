SET @dbname = "fitness";
SET @tablename = "workout_sets";
SET @key = "fk_workout_sets_action_id_to_actions_id";

SET @dropForeignKey = (SELECT IF(
  (
    SELECT COUNT(*) FROM information_schema.REFERENTIAL_CONSTRAINTS
    WHERE
      (TABLE_NAME = @tablename)
      AND (CONSTRAINT_SCHEMA = @dbname)
      AND (CONSTRAINT_NAME = "fk_workout_sets_action_id_to_actions_id")
  ) > 0,
    CONCAT("ALTER TABLE ", @tablename, " DROP FOREIGN KEY `fk_workout_sets_action_id_to_actions_id`"),
    "SELECT 1"
));

SET @addForeignKey = CONCAT("ALTER TABLE ", @tablename, " ADD CONSTRAINT ", @key, " FOREIGN KEY (`action_id`) REFERENCES `actions` (`id`) ON DELETE SET NULL");

PREPARE dropForeignKey FROM @dropForeignKey;
EXECUTE dropForeignKey;
DEALLOCATE PREPARE dropForeignKey;

PREPARE addForeignKey FROM @addForeignKey;
EXECUTE addForeignKey;
DEALLOCATE PREPARE addForeignKey;