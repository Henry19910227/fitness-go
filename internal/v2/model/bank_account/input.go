package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/bank_account/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/bank_account/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	optional.UserIDField
	PagingInput
	OrderByInput
}

type FindInput struct {
	optional.UserIDField
}

// APIGetTrainerBankAccountInput /v2/trainer/bank_account [GET]
type APIGetTrainerBankAccountInput struct {
	required.UserIDField
}

// APIUpdateTrainerBankAccountInput /v2/trainer/bank_account [PATCH]
type APIUpdateTrainerBankAccountInput struct {
	required.UserIDField
	Form APIUpdateTrainerBankAccountForm
}
type APIUpdateTrainerBankAccountForm struct {
	AccountImageFile *file.Input
	optional.AccountNameField
	optional.BankCodeField
	optional.BranchField
	optional.AccountField
}
