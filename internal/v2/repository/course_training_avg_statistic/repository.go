package course_training_avg_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 course_id 篩選條件
	if input.CourseID != nil {
		db = db.Where("course_id = ?", *input.CourseID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	//加入 course_id 篩選條件
	if input.CourseID != nil {
		db = db.Where("course_training_avg_statistics.course_id = ?", *input.CourseID)
	}
	// Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("course_training_avg_statistics.*")
	// Paging
	if input.Page != nil && input.Size != nil {
		db = db.Offset((*input.Page - 1) * *input.Size).Limit(*input.Size)
	} else if input.Page != nil {
		db = db.Offset(0)
	} else if input.Size != nil {
		db = db.Limit(*input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("course_training_avg_statistics.%s %s", input.OrderField, input.OrderType))
	}
	// Custom Order
	if input.Orders != nil {
		for _, orderBy := range input.Orders {
			db = db.Order(orderBy.Value)
		}
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

// Statistic SQL
/*
INSERT INTO course_training_avg_statistics (course_id, rate)
SELECT a.id AS course_id, TRUNCATE(AVG(a.total_workout_count) / courses.workout_count * 100, 0) AS rate
FROM courses
INNER JOIN (
	SELECT b.id,b.user_id,COUNT(*) AS total_workout_count
	FROM (
		SELECT courses.id,workout_logs.user_id,workout_logs.workout_id,COUNT(*) AS total_workout_logs
		FROM courses
		INNER JOIN plans ON courses.id = plans.course_id
		INNER JOIN workouts ON plans.id = workouts.plan_id
		INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
		GROUP BY courses.id,workout_logs.user_id,workout_logs.workout_id
	) AS b
    GROUP BY b.id,b.user_id
) AS a ON a.id = courses.id
WHERE courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3
GROUP BY a.id
ON DUPLICATE KEY UPDATE rate = VALUES(rate), update_at = CURRENT_TIMESTAMP;
*/
func (r *repository) Statistic() (err error) {
	err = r.db.Exec("INSERT INTO course_training_avg_statistics (course_id, rate)\n" +
		"SELECT a.id AS course_id, TRUNCATE(AVG(a.total_workout_count) / courses.workout_count * 100, 0) AS rate\n" +
		"FROM courses\n" +
		"INNER JOIN (\n" +
		"\tSELECT b.id,b.user_id,COUNT(*) AS total_workout_count\n" +
		"\tFROM (\n" +
		"\t\tSELECT courses.id,workout_logs.user_id,workout_logs.workout_id,COUNT(*) AS total_workout_logs\n" +
		"\t\tFROM courses\n" +
		"\t\tINNER JOIN plans ON courses.id = plans.course_id\n" +
		"\t\tINNER JOIN workouts ON plans.id = workouts.plan_id\n" +
		"\t\tINNER JOIN workout_logs ON workouts.id = workout_logs.workout_id\n" +
		"\t\tGROUP BY courses.id,workout_logs.user_id,workout_logs.workout_id\n" +
		"\t) AS b\n" +
		"\tGROUP BY b.id,b.user_id\n" +
		") AS a ON a.id = courses.id\n" +
		"WHERE courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3\n" +
		"GROUP BY a.id\n" +
		"ON DUPLICATE KEY UPDATE rate = VALUES(rate), update_at = CURRENT_TIMESTAMP").Error
	return err
}
