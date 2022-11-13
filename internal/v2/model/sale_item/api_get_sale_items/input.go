package api_get_sale_items

import (
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type PagingInput = paging.Input

// Input /v2/sale_items [GET]
type Input struct {
	userRequired.UserIDField
}
