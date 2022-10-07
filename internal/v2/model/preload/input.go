package preload

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
)

type Preload struct {
	Field      string
	Conditions []interface{}
	OrderBy    *orderBy.Input
}

type Input struct {
	Preloads []*Preload
}
