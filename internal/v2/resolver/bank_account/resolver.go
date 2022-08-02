package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
)

type resolver struct {
	bankAccountService bank_account.Service
}

func New(bankAccountService bank_account.Service) Resolver {
	return &resolver{bankAccountService: bankAccountService}
}

func (r *resolver) APIGetTrainerBankAccount(input *model.APIGetTrainerBankAccountInput) (output model.APIGetTrainerBankAccountOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Page = 1
	listInput.Size = 1
	datas, _, err := r.bankAccountService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(datas) == 0 {
		output.Set(code.DataNotFound, "查無資料")
		return output
	}
	// parser output
	data := model.APIGetTrainerBankAccountData{}
	if err := util.Parser(datas[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
