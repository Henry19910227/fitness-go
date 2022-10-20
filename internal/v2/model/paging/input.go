package paging

import "github.com/Henry19910227/fitness-go/internal/v2/field/paging/required"

type Input struct {
	required.PageField
	required.SizeField
}
