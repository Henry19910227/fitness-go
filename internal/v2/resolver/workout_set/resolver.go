package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	actionModel "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	workoutModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/service/action"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"gorm.io/gorm"
)

type resolver struct {
	workoutSetService workout_set.Service
	workoutService    workout.Service
	courseService     course.Service
	actionService     action.Service
	startAudioTool    uploader.Tool
	ProgressAudioTool uploader.Tool
}

func New(workoutSetService workout_set.Service, workoutService workout.Service,
	courseService course.Service, actionService action.Service,
	startAudioTool uploader.Tool, ProgressAudioTool uploader.Tool) Resolver {
	return &resolver{workoutSetService: workoutSetService, workoutService: workoutService,
		courseService: courseService, actionService: actionService,
		startAudioTool: startAudioTool, ProgressAudioTool: ProgressAudioTool}
}

func (r *resolver) APICreateUserWorkoutSets(tx *gorm.DB, input *model.APICreateUserWorkoutSetsInput) (output model.APICreateUserWorkoutSetsOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法創建資源")
		return output
	}
	// 驗證動作權限
	actionListInput := actionModel.ListInput{}
	actionListInput.IDs = input.Body.ActionIDs
	actionListInput.Type = util.PointerInt(1)
	actionOutputs, _, err := r.actionService.Tx(tx).List(&actionListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(actionOutputs) != len(input.Body.ActionIDs) {
		output.Set(code.BadRequest, "含有不合法動作")
		return output
	}
	for _, actionOutput := range actionOutputs {
		if util.OnNilJustReturnInt(actionOutput.Source, 0) == actionModel.SourceTrainer {
			output.Set(code.PermissionDenied, "不可添加教練動作")
			return output
		}
	}
	// 添加訓練組
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

func (r *resolver) APICreateUserWorkoutSetByDuplicate(tx *gorm.DB, input *model.APICreateUserWorkoutSetByDuplicateInput) (output model.APICreateUserWorkoutSetByDuplicateOutput) {
	defer tx.Rollback()
	/** 驗證權限流程 */
	// 1.查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutSetID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法創建資源")
		return output
	}
	/** 複製訓練組流程 */
	// 1.查詢訓練組
	findWorkoutSetInput := model.FindInput{}
	findWorkoutSetInput.ID = util.PointerInt64(input.Uri.ID)
	workoutSetOutput, err := r.workoutSetService.Tx(tx).Find(&findWorkoutSetInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.複製訓練組
	workoutSetTables := make([]*model.Table, 0)
	for i := 0; i < input.Body.DuplicateCount; i++ {
		workoutSetTable := model.Table{}
		if err := util.Parser(workoutSetOutput, &workoutSetTable); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		workoutSetTable.ID = nil
		workoutSetTables = append(workoutSetTables, &workoutSetTable)
	}
	// 3.新增訓練組
	_, err = r.workoutSetService.Tx(tx).Create(workoutSetTables)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	/** 更新上層訓練的訓練組數量流程 */
	// 1.查詢當前訓練組數量
	setListInput := model.ListInput{}
	setListInput.WorkoutID = workoutSetOutput.WorkoutID
	setListInput.Type = util.PointerInt(1)
	setOutputs, _, err := r.workoutSetService.Tx(tx).List(&setListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2.更新當前訓練組數量
	workoutTable := workoutModel.Table{}
	workoutTable.ID = workoutSetOutput.WorkoutID
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

func (r *resolver) APICreateUserRestSet(input *model.APICreateUserRestSetInput) (output model.APICreateUserRestSetOutput) {
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	courseOutput, err := r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法創建資源")
		return output
	}
	// 添加休息組
	table := model.Table{}
	table.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	table.Type = util.PointerInt(2)
	table.AutoNext = util.PointerString("N")
	table.Weight = util.PointerFloat64(0)
	table.Reps = util.PointerInt(0)
	table.Distance = util.PointerFloat64(0)
	table.Duration = util.PointerInt(30)
	table.Incline = util.PointerFloat64(0)
	table.StartAudio = util.PointerString("")
	table.ProgressAudio = util.PointerString("")
	table.Remark = util.PointerString("")
	_, err = r.workoutSetService.Create([]*model.Table{&table})
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteUserWorkoutSet(tx *gorm.DB, input *model.APIDeleteUserWorkoutSetInput) (output model.APIDeleteUserWorkoutSetOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutSetID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢訓練
	findSetInput := model.FindInput{}
	findSetInput.ID = util.PointerInt64(input.Uri.ID)
	workoutSetOutput, err := r.workoutSetService.Tx(tx).Find(&findSetInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法刪除資源")
		return output
	}
	// 刪除訓練組
	deleteSetInput := model.DeleteInput{}
	deleteSetInput.ID = input.Uri.ID
	if err := r.workoutSetService.Tx(tx).Delete(&deleteSetInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此訓練的訓練組數量
	setListInput := model.ListInput{}
	setListInput.WorkoutID = workoutSetOutput.WorkoutID
	setListInput.Type = util.PointerInt(1)
	setOutputs, _, err := r.workoutSetService.Tx(tx).List(&setListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此訓練的訓練組數量
	workoutTable := workoutModel.Table{}
	workoutTable.ID = workoutSetOutput.WorkoutID
	workoutTable.WorkoutSetCount = util.PointerInt(len(setOutputs))
	if err := r.workoutService.Tx(tx).Update(&workoutTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除 start audio 檔案
	_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutSetOutput.StartAudio, ""))
	// 刪除 progress audio 檔案
	_ = r.ProgressAudioTool.Delete(util.OnNilJustReturnString(workoutSetOutput.ProgressAudio, ""))

	tx.Commit()
	// Parser Output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetUserWorkoutSets(input *model.APIGetUserWorkoutSetsInput) (output model.APIGetUserWorkoutSetsOutput) {
	// 查詢
	listInput := model.ListInput{}
	listInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Action"},
	}
	setOutputs, _, err := r.workoutSetService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserWorkoutSetsData{}
	if err := util.Parser(setOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}

func (r *resolver) APIUpdateUserWorkoutSet(tx *gorm.DB, input *model.APIUpdateUserWorkoutSetInput) (output model.APIUpdateUserWorkoutSetOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutSetID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此訓練組擁有者，無法修改資源")
		return output
	}
	// 查詢訓練組資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	workoutSetOutput, err := r.workoutSetService.Tx(tx).Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練組
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.workoutSetService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 上傳 start audio
	if input.Form.StartAudio != nil {
		// 儲存 start audio
		startAudioNamed, err := r.startAudioTool.Save(input.Form.StartAudio.Data, input.Form.StartAudio.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練組
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.StartAudio = util.PointerString(startAudioNamed)
		if err := r.workoutSetService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊的 start audio
		_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutSetOutput.StartAudio, ""))
	}
	// 上傳 progress audio
	if input.Form.ProgressAudio != nil {
		// 儲存 progress audio
		endAudioNamed, err := r.ProgressAudioTool.Save(input.Form.ProgressAudio.Data, input.Form.ProgressAudio.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練組
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.ProgressAudio = util.PointerString(endAudioNamed)
		if err := r.workoutSetService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊的 progress audio
		_ = r.ProgressAudioTool.Delete(util.OnNilJustReturnString(workoutSetOutput.ProgressAudio, ""))
	}
	tx.Commit()
	// parser output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteUserWorkoutSetStartAudio(input *model.APIDeleteUserWorkoutSetStartAudioInput) (output model.APIDeleteUserWorkoutSetStartAudioOutput) {
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutSetID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法刪除資源")
		return output
	}
	// 查詢訓練資訊
	findWorkoutSetInput := model.FindInput{}
	findWorkoutSetInput.ID = util.PointerInt64(input.Uri.ID)
	workoutSetOutput, err := r.workoutSetService.Find(&findWorkoutSetInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練組
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.StartAudio = util.PointerString("")
	if err := r.workoutSetService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除音檔
	_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutSetOutput.StartAudio, ""))
	// parser output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteUserWorkoutSetProgressAudio(input *model.APIDeleteUserWorkoutSetProgressAudioInput) (output model.APIDeleteUserWorkoutSetProgressAudioOutput) {
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutSetID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法刪除資源")
		return output
	}
	// 查詢訓練資訊
	findWorkoutSetInput := model.FindInput{}
	findWorkoutSetInput.ID = util.PointerInt64(input.Uri.ID)
	workoutSetOutput, err := r.workoutSetService.Find(&findWorkoutSetInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練組
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.ProgressAudio = util.PointerString("")
	if err := r.workoutSetService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除音檔
	_ = r.ProgressAudioTool.Delete(util.OnNilJustReturnString(workoutSetOutput.ProgressAudio, ""))
	// parser output
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
