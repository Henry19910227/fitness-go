package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "bank_accounts"
}

// APIGetTrainerBankAccountOutput /v2/trainer/bank_account [GET]
type APIGetTrainerBankAccountOutput struct {
	base.Output
	Data *APIGetTrainerBankAccountData `json:"data,omitempty"`
}
type APIGetTrainerBankAccountData struct {
	AccountNameField
	AccountImageField
	BackCodeField
	BranchField
	AccountField
	CreateAtField
	UpdateAtField
}

// APIUpdateTrainerBankAccountOutput /v2/trainer/bank_account [GET]
type APIUpdateTrainerBankAccountOutput struct {
	base.Output
}
