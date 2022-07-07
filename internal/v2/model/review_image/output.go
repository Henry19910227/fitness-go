package review_image

type Output struct {
	Table
}

func (Output) TableName() string {
	return "review_images"
}
