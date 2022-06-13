package model

import "github.com/Henry19910227/fitness-go/internal/pkg/global"

type PagingParam struct {
	Offset int
	Limit  int
}

type OrderBy struct {
	Field     string
	OrderType global.OrderType
}

type Preload struct {
	Field   string
}

type InPageParam struct {
	Paging *PagingParam
}

type OrderByParam struct {
	OrderBy *OrderBy
}

type PreloadParam struct {
	Preloads []*Preload
}

type DeletedParam struct {
	IsDeleted *int
}
