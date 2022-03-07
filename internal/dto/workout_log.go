package dto

type WorkoutLogResponse struct {
	Duration            int                  `json:"duration" example:"1"`
	Intensity           int                  `json:"intensity" example:"1"`
	Place               int                  `json:"place" example:"1"`
	UserCourseStatistic *UserCourseStatistic `json:"user_course_statistic"`
	WorkoutSetLogs      []*WorkoutSetLog     `json:"workout_set_logs"`
}

type CreateWorkoutLogParam struct {
	Duration       int
	Intensity      int
	Place          int
	WorkoutSetLogs []*WorkoutSetLog
}
