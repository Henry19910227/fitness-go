package rda

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/rda"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_update_rda"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/rda"
	"gorm.io/gorm"
	"time"
)

type resolver struct {
	rdaService  rda.Service
	dietService diet.Service
}

func New(rdaService rda.Service, dietService diet.Service) Resolver {
	return &resolver{rdaService: rdaService, dietService: dietService}
}

func (r *resolver) APIUpdateRDA(tx *gorm.DB, input *api_update_rda.Input) (output api_update_rda.Output) {
	defer tx.Rollback()
	// 創建 rda
	rdaTable := model.Table{}
	rdaTable.UserID = util.PointerInt64(input.UserID)
	if err := util.Parser(input.Body, &rdaTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	rdaID, err := r.rdaService.Tx(tx).Create(&rdaTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查找今天之後的飲食紀錄
	dietListInput := dietModel.ListInput{}
	dietListInput.UserID = util.PointerInt64(input.UserID)
	dietListInput.Wheres = []*whereModel.Where{
		{Query: "DATE_FORMAT(diets.schedule_at,'%Y-%m-%d') >= DATE_FORMAT(?,'%Y-%m-%d')", Args: []interface{}{util.PointerString(time.Now().Format("2006-01-02"))}},
	}
	dietOutputs, _, err := r.dietService.Tx(tx).List(&dietListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新今天之後的飲食紀錄
	dietTables := make([]*dietModel.Table, 0)
	for _, dietOutput := range dietOutputs {
		dietTable := dietModel.Table{}
		dietTable.ID = dietOutput.ID
		dietTable.RdaID = util.PointerInt64(rdaID)
		dietTables = append(dietTables, &dietTable)
	}
	if err := r.dietService.Tx(tx).Updates(dietTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}
