package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
)

type resolver struct {
	bankAccountService bank_account.Service
	uploadTool         uploader.Tool
}

func New(bankAccountService bank_account.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{bankAccountService: bankAccountService, uploadTool: uploadTool}
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

func (r *resolver) APIUpdateTrainerBankAccount(input *model.APIUpdateTrainerBankAccountInput) (output model.APIUpdateTrainerBankAccountOutput) {
	//查詢當前資料
	findInput := model.FindInput{}
	findInput.UserID = util.PointerInt64(input.UserID)
	data, err := r.bankAccountService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser input
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//存取圖片
	if input.Form.AccountImageFile != nil {
		imageNamed, err := r.uploadTool.Save(input.Form.AccountImageFile.Data, input.Form.AccountImageFile.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.AccountImage = util.PointerString(imageNamed)
	}
	//執行更新
	if err := r.bankAccountService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除舊圖片
	_ = r.uploadTool.Delete(util.OnNilJustReturnString(data.AccountImage, ""))
	output.Set(code.Success, "success")
	return output
}
