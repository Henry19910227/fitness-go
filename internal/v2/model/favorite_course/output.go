package favorite_course

type Output struct {
	Table
}

func (Output) TableName() string {
	return "favorite_courses"
}