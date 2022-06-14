package order_by

func NewInput(orderField string, orderType string) *Input {
	i := Input{}
	i.OrderField = orderField
	i.OrderType = orderType
	return &i
}
