package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	bodyService "github.com/Henry19910227/fitness-go/internal/v2/service/body_record"
)

type resolver struct {
	bodyService bodyService.Service
}

func New(bodyService bodyService.Service) Resolver {
	return &resolver{bodyService: bodyService}
}

func (r *resolver) APICreateBodyRecord(input *model.APICreateBodyRecordInput) (output model.APICreateBodyRecordOutput) {
	table := model.Table{}
	table.UserID = input.UserID
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	result, err := r.bodyService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateBodyRecordData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
