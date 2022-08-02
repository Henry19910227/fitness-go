package bank_account

import model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"

type Resolver interface {
	APIGetTrainerBankAccount(input *model.APIGetTrainerBankAccountInput) (output model.APIGetTrainerBankAccountOutput)
	APIUpdateTrainerBankAccount(input *model.APIUpdateTrainerBankAccountInput) (output model.APIUpdateTrainerBankAccountOutput)
}
