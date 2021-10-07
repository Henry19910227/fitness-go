package model

import "github.com/Henry19910227/fitness-go/internal/global"

type PagingParam struct {
	Offset int
	Limit int
}

type OrderBy struct {
	Field string
	OrderType global.OrderType
}
