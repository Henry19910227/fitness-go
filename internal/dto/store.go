package dto

type StoreHomePage struct {
	PopularCourses []*CourseSummary `json:"popular_courses"` // 熱門課表
	PopularTrainers []*TrainerSummary `json:"popular_trainers"` // 熱門教練
	LatestCourses []*CourseSummary `json:"latest_courses"` // 最新課表
	LatestTrainers []*TrainerSummary `json:"latest_trainers"` // 最新教練
}
