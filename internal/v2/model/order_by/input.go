package order_by

const (
	ASC  string = "ASC"
	DESC string = "DESC"
)

type Order struct {
	Query string
	Args []interface{}
}

type Input struct {
	OrderFieldField
	OrderTypeField
}

type CustomInput struct {
	Orders []*Order
}
