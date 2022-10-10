package user_course_asset

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_course_assets"
}
