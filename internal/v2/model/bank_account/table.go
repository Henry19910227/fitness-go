package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/bank_account/optional"
)

type 	Table struct {
	optional.UserIDField
	optional.AccountNameField
	optional.AccountField
	optional.AccountImageField
	optional.BankCodeField
	optional.BranchField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "bank_accounts"
}
