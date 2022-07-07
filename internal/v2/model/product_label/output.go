package product_label

type Output struct {
	Table
}

func (Output) TableName() string {
	return "product_labels"
}
