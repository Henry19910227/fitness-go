package course_usage_statistic

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
}

// Statistic 完整SQL https://kind-bass-788.notion.site/course_usage_statistic-7976a7b0b9cc4dc3b10fa06819be1f44
func (r *repository) Statistic() (err error) {
	err = r.db.Exec("INSERT INTO course_usage_statistics " +
		"( " +
		"course_id, " +
		"total_finish_workout_count, " +
		"user_finish_count, " +
		"male_finish_count, " +
		"female_finish_count, " +
		"finish_count_avg, " +
		"age_13_17_count, " +
		"age_18_24_count, " +
		"age_25_34_count, " +
		"age_35_44_count, " +
		"age_45_54_count, " +
		"age_55_64_count, " +
		"age_65_up_count " +
		") " +
		"SELECT " +
		"IFNULL(a.course_id, 0) AS course_id, " +
		"IFNULL(a.total_finish_workout_count, 0) AS total_finish_workout_count, " +
		"IFNULL(b.user_finish_count, 0) AS user_finish_count, " +
		"IFNULL(c.male_finish_count, 0) AS male_finish_count, " +
		"IFNULL(d.female_finish_count, 0) AS female_finish_count ," +
		"IFNULL(e.finish_count_avg, 0) AS finish_count_avg ," +
		"IFNULL(f.age_13_17_count, 0) AS age_13_17_count ," +
		"IFNULL(g.age_18_24_count, 0) AS age_18_24_count ," +
		"IFNULL(h.age_25_34_count, 0) AS age_25_34_count ," +
		"IFNULL(i.age_35_44_count, 0) AS age_35_44_count ," +
		"IFNULL(j.age_45_54_count, 0) AS age_45_54_count ," +
		"IFNULL(k.age_55_64_count, 0) AS age_55_64_count ," +
		"IFNULL(l.age_65_up_count, 0) AS age_65_up_count " +
		"FROM " +
		"( " + tableA() + " ) AS a " +
		"LEFT JOIN " +
		"( " + tableB() + " ) AS b ON a.course_id = b.course_id " +
		"LEFT JOIN " +
		"( " + tableC() + " ) AS c ON a.course_id = c.course_id " +
		"LEFT JOIN " +
		"( " + tableD() + " ) AS d ON a.course_id = d.course_id " +
		"LEFT JOIN " +
		"( " + tableE() + " ) AS e ON a.course_id = e.course_id " +
		"LEFT JOIN " +
		"( " + tableF() + " ) AS f ON a.course_id = f.course_id " +
		"LEFT JOIN " +
		"( " + tableG() + " ) AS g ON a.course_id = g.course_id " +
		"LEFT JOIN " +
		"( " + tableH() + " ) AS h ON a.course_id = h.course_id " +
		"LEFT JOIN " +
		"( " + tableI() + " ) AS i ON a.course_id = i.course_id " +
		"LEFT JOIN " +
		"( " + tableJ() + " ) AS j ON a.course_id = j.course_id " +
		"LEFT JOIN " +
		"( " + tableK() + " ) AS k ON a.course_id = k.course_id " +
		"LEFT JOIN " +
		"( " + tableL() + " ) AS l ON a.course_id = l.course_id " +
		"ON DUPLICATE KEY UPDATE " +
		"total_finish_workout_count = IFNULL(a.total_finish_workout_count, 0), " +
		"user_finish_count = IFNULL(b.user_finish_count, 0), " +
		"male_finish_count = IFNULL(c.male_finish_count, 0), " +
		"female_finish_count = IFNULL(d.female_finish_count, 0), " +
		"finish_count_avg = IFNULL(e.finish_count_avg, 0), " +
		"age_13_17_count = IFNULL(f.age_13_17_count, 0), " +
		"age_18_24_count = IFNULL(g.age_18_24_count, 0), " +
		"age_25_34_count = IFNULL(h.age_25_34_count, 0), " +
		"age_35_44_count = IFNULL(i.age_35_44_count, 0), " +
		"age_45_54_count = IFNULL(j.age_45_54_count, 0), " +
		"age_55_64_count = IFNULL(k.age_55_64_count, 0), " +
		"age_65_up_count = IFNULL(l.age_65_up_count, 0)").Error
	return err
}

func tableA() string {
	return "SELECT MAX(course_id) AS course_id, SUM(total_finish_workout_count) AS total_finish_workout_count " +
		"FROM user_course_statistics " +
		"GROUP BY user_course_statistics.course_id"
}
func tableB() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS user_finish_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"GROUP BY t.course_id"
}

func tableC() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"INNER JOIN users ON workout_logs.user_id = users.id " +
		"WHERE users.sex = 'm' " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS male_finish_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"GROUP BY t.course_id"
}

func tableD() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"INNER JOIN users ON workout_logs.user_id = users.id " +
		"WHERE users.sex = 'f' " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS female_finish_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"GROUP BY t.course_id"
}

func tableE() string {
	return "SELECT MAX(course_id) AS course_id, ROUND(AVG(finish_workout_count)) AS finish_count_avg " +
		"FROM user_course_statistics " +
		"GROUP BY user_course_statistics.course_id"
}

func tableF() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_13_17_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '13' AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) < '17' " +
		"GROUP BY t.course_id"
}

func tableG() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_18_24_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '18' AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) < '24' " +
		"GROUP BY t.course_id"
}

func tableH() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_25_34_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '25' AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) < '34' " +
		"GROUP BY t.course_id"
}

func tableI() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_35_44_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '35' AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) < '44' " +
		"GROUP BY t.course_id"
}

func tableJ() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_45_54_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '45' AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) < '54' " +
		"GROUP BY t.course_id"
}

func tableK() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_55_64_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '55' AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) < '64' " +
		"GROUP BY t.course_id"
}

func tableL() string {
	t := "SELECT courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM courses " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT t.course_id AS course_id, COUNT(*) AS age_65_up_count " +
		"FROM " +
		"( " + t + " ) AS t " +
		"INNER JOIN users ON t.user_id = users.id " +
		"WHERE TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= '65' " +
		"GROUP BY t.course_id"
}
