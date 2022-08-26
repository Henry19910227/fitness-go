package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	actionModel "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	workoutModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/service/action"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"gorm.io/gorm"
)

type resolver struct {
	workoutSetService workout_set.Service
	workoutService workout.Service
	actionService action.Service
}

func New(workoutSetService workout_set.Service, workoutService workout.Service, actionService action.Service) Resolver {
	return &resolver{workoutSetService: workoutSetService, workoutService: workoutService, actionService: actionService}
}

func (r *resolver) APICreateUserWorkoutSets(tx *gorm.DB, input *model.APICreateUserWorkoutSetsInput) (output model.APICreateUserWorkoutSetsOutput) {
	defer tx.Rollback()
	// 驗證動作權限
	actionListInput := actionModel.ListInput{}
	actionListInput.IDs = input.Body.ActionIDs
	actionListInput.Type = util.PointerInt(1)
	actionOutputs, _, err := r.actionService.List(&actionListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(actionOutputs) != len(input.Body.ActionIDs) {
		output.Set(code.BadRequest, "含有不合法動作")
		return output
	}
	for _, actionOutput := range actionOutputs{
		if util.OnNilJustReturnInt(actionOutput.Source, 0) == actionModel.SourceTrainer {
			output.Set(code.PermissionDenied, "不可添加教練動作")
			return output
		}
	}
	// 添加動作組
	setTables := make([]*model.Table, 0)
	for _, actionID := range input.Body.ActionIDs {
		setTable := model.Table{}
		setTable.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
		setTable.ActionID = util.PointerInt64(actionID)
		setTable.Type = util.PointerInt(1)
		setTable.AutoNext = util.PointerString("N")
		setTable.Weight = util.PointerFloat64(10)
		setTable.Reps = util.PointerInt(10)
		setTable.Distance = util.PointerFloat64(1)
		setTable.Duration = util.PointerInt(60)
		setTable.Incline = util.PointerFloat64(1)
		setTable.StartAudio = util.PointerString("")
		setTable.ProgressAudio = util.PointerString("")
		setTable.Remark = util.PointerString("")
		setTables = append(setTables, &setTable)
	}
	_, err = r.workoutSetService.Tx(tx).Create(setTables)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此訓練的訓練組數量
	setListInput := model.ListInput{}
	setListInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	setListInput.Type = util.PointerInt(1)
	setOutputs, _, err := r.workoutSetService.Tx(tx).List(&setListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此訓練的訓練組數量
	workoutTable := workoutModel.Table{}
	workoutTable.ID = util.PointerInt64(input.Uri.WorkoutID)
	workoutTable.WorkoutSetCount = util.PointerInt(len(setOutputs))
	if err := r.workoutService.Tx(tx).Update(&workoutTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetCMSWorkoutSets(input *model.APIGetCMSWorkoutSetsInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Action"},
	}
	// 調用 repo
	result, page, err := r.workoutSetService.List(&param)
	if err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSWorkoutSetsData{}
	if err := util.Parser(result, &data); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSWorkoutSetsOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}
