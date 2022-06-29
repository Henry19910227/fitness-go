package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
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
	table.UserID = util.PointerInt64(input.UserID)
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

func (r *resolver) APIGetBodyRecords(input *model.APIGetBodyRecordsInput) (output model.APIGetBodyRecordsOutput) {
	bodyInput := model.ListInput{}
	bodyInput.UserID = util.PointerInt64(input.UserID)
	bodyInput.OrderField = "create_at"
	bodyInput.OrderType = order_by.DESC
	if err := util.Parser(input.Query, &bodyInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	datas, page, err := r.bodyService.List(&bodyInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetBodyRecordsData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetBodyRecordsLatest(input *model.APIGetBodyRecordsLatestInput) (output model.APIGetBodyRecordsLatestOutput) {
	// parser input
	listInput := model.LatestListInput{}
	listInput.UserID = input.UserID
	// LatestList
	listData, err := r.bodyService.LatestList(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetBodyRecordsLatestData{}
	if err := util.Parser(listData, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}

func (r *resolver) APIUpdateBodyRecord(input *model.APIUpdateBodyRecordInput) (output base.Output) {
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.Value = input.Body.Value
	// 更新資料
	if err := r.bodyService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteBodyRecord(input *model.APIDeleteBodyRecordInput) (output base.Output) {
	deleteInput := model.DeleteInput{}
	deleteInput.ID = util.PointerInt64(input.Uri.ID)
	if err := r.bodyService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}
