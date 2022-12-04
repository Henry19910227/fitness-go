package paging

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/paging/required"
)

type Output struct {
	required.TotalCountField
	optional.TotalPageField
	optional.PageField
	optional.SizeField
}
