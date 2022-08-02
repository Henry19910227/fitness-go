package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
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

type FindInput struct {
	UserIDOptional
}

// APIGetTrainerBankAccountInput /v2/trainer/bank_account [GET]
type APIGetTrainerBankAccountInput struct {
	UserIDRequired
}

// APIUpdateTrainerBankAccountInput /v2/trainer/bank_account [PATCH]
type APIUpdateTrainerBankAccountInput struct {
	UserIDRequired
	Form APIUpdateTrainerBankAccountForm
}
type APIUpdateTrainerBankAccountForm struct {
	AccountImageFile *file.Input
	AccountNameOptional
	BankCodeOptional
	BranchOptional
	AccountOptional
}
