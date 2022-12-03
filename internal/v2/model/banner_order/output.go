package banner_order

type Output struct {
	Table
}

func (Output) TableName() string {
	return "banner_orders"
}
