package order_by

const (
	ASC  string = "ASC"
	DESC string = "DESC"
)

type Order struct {
	Value interface{}
}

type Input struct {
	OrderFieldField
	OrderTypeField
}

type CustomInput struct {
	Orders []*Order
}
