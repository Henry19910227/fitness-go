package feedback_image

type Output struct {
	Table
}

func (Output) TableName() string {
	return "feedback_images"
}
