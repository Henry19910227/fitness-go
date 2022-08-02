package bank_account

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	UserIDOptional
	PagingInput
	OrderByInput
}

// APIGetTrainerBankAccountInput /v2/trainer/bank_account [GET]
type APIGetTrainerBankAccountInput struct {
	UserIDRequired
}
