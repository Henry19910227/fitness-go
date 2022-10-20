package paging

import "github.com/Henry19910227/fitness-go/internal/v2/field/paging/required"

type Output struct {
	required.TotalCountField
	required.TotalPageField
	required.PageField
	required.SizeField
}
