package order_by

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/order_by/required"
)

const (
	ASC  string = "ASC"
	DESC string = "DESC"
)

type Order struct {
	Value interface{}
}

type Input struct {
	required.OrderFieldField
	required.OrderTypeField
}

type CustomInput struct {
	Orders []*Order
}
