package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	planModel "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	workoutSetModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"gorm.io/gorm"
)

type resolver struct {
	workoutService    workout.Service
	workoutSetService workout_set.Service
	planService       plan.Service
	courseService     course.Service
	startAudioTool    uploader.Tool
	endAudioTool      uploader.Tool
}

func New(workoutService workout.Service, workoutSetService workout_set.Service, planService plan.Service, courseService course.Service, startAudioTool uploader.Tool, endAudioTool uploader.Tool) Resolver {
	return &resolver{workoutService: workoutService, workoutSetService: workoutSetService, planService: planService, courseService: courseService, startAudioTool: startAudioTool, endAudioTool: endAudioTool}
}

func (r *resolver) APICreateUserWorkout(tx *gorm.DB, input *model.APICreateUserWorkoutInput) (output model.APICreateUserWorkoutOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證課表權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.WorkoutCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達計畫數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法創建資源")
		return output
	}
	// 創建訓練
	workoutTable := model.Table{}
	workoutTable.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutTable.Name = util.PointerString(input.Body.Name)
	workoutTable.Equipment = util.PointerString("")
	workoutTable.StartAudio = util.PointerString("")
	workoutTable.EndAudio = util.PointerString("")
	workoutTable.WorkoutSetCount = util.PointerInt(0)
	workoutID, err := r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此計畫訓練數量
	workoutListInput := model.ListInput{}
	workoutListInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此計畫訓練數量
	planTable := planModel.Table{}
	planTable.ID = util.PointerInt64(input.Uri.PlanID)
	planTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.planService.Tx(tx).Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表訓練數量
	workoutListInput = model.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err = r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此課表訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateUserWorkoutData{}
	data.ID = util.PointerInt64(workoutID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APICreateUserWorkoutFromTemplate(tx *gorm.DB, input *model.APICreateUserWorkoutInput) (output model.APICreateUserWorkoutOutput) {
	defer tx.Rollback()
	/** 驗證權限流程 */
	// 1.驗證課表權限
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.WorkoutCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達計畫數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法創建資源")
		return output
	}
	// 2.驗證模板 workout_id 合法性
	findCourseInput = courseModel.FindInput{}
	findCourseInput.WorkoutID = input.Body.WorkoutTemplateID
	templateOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.ID, 0) != util.OnNilJustReturnInt64(templateOutput.ID, 0) {
		output.Set(code.BadRequest, "非此課表所屬模板id，無法創建資源")
		return output
	}
	/** 查詢模板資訊 */
	// 查詢模板 workout 資訊
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = input.Body.WorkoutTemplateID
	workoutOutput, err := r.workoutService.Tx(tx).Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢模板 workout_set 資訊
	workoutSetListInput := workoutSetModel.ListInput{}
	workoutSetListInput.WorkoutID = input.Body.WorkoutTemplateID
	workoutSetOutputs, _, err := r.workoutSetService.Tx(tx).List(&workoutSetListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	/** 以模板創建 workout */
	// 1. 創建 workout 到指定 plan_id
	workoutTable := model.Table{}
	workoutTable.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutTable.Name = util.PointerString(input.Body.Name)
	workoutTable.Equipment = workoutOutput.Equipment
	workoutTable.StartAudio = workoutOutput.StartAudio
	workoutTable.EndAudio = workoutOutput.EndAudio
	workoutTable.WorkoutSetCount = workoutOutput.WorkoutSetCount
	newWorkoutID, err := r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2. 創建 workout_sets
	workoutSetTables := make([]*workoutSetModel.Table, 0)
	if err := util.Parser(workoutSetOutputs, &workoutSetTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, workoutSetTable := range workoutSetTables {
		workoutSetTable.ID = nil
		workoutSetTable.WorkoutID = util.PointerInt64(newWorkoutID)
	}
	_, err = r.workoutSetService.Tx(tx).Create(workoutSetTables)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	/** 更新訓練與訓練組數 */
	// 1. 查詢此課表訓練數量
	workoutListInput := model.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2. 更新此課表訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3. 查詢此計畫訓練數量
	workoutListInput = model.ListInput{}
	workoutListInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutOutputs, _, err = r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4. 更新此計畫訓練數量
	planTable := planModel.Table{}
	planTable.ID = util.PointerInt64(input.Uri.PlanID)
	planTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.planService.Tx(tx).Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 5. 更新此訓練的訓練組數量
	workoutTable = model.Table{}
	workoutTable.ID = util.PointerInt64(newWorkoutID)
	workoutTable.WorkoutSetCount = util.PointerInt(len(workoutSetOutputs))
	if err := r.workoutService.Tx(tx).Update(&workoutTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateUserWorkoutData{}
	data.ID = util.PointerInt64(newWorkoutID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIDeleteUserWorkout(tx *gorm.DB, input *model.APIDeleteUserWorkoutInput) (output model.APIDeleteUserWorkoutOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢訓練
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err := r.workoutService.Tx(tx).Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證課表刪除權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout {
		output.Set(code.BadRequest, "單一訓練課表，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	// 刪除訓練
	deleteWorkoutInput := model.DeleteInput{}
	deleteWorkoutInput.ID = input.Uri.ID
	if err := r.workoutService.Tx(tx).Delete(&deleteWorkoutInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表訓練數量
	workoutListInput := model.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此課表訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此計畫訓練數量
	workoutListInput = model.ListInput{}
	workoutListInput.PlanID = workoutOutput.PlanID
	workoutOutputs, _, err = r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此計畫訓練數量
	planTable := planModel.Table{}
	planTable.ID = workoutOutput.PlanID
	planTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.planService.Tx(tx).Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除 start audio 檔案
	_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.StartAudio, ""))
	// 刪除 progress audio 檔案
	_ = r.endAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.EndAudio, ""))

	tx.Commit()
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetUserWorkouts(input *model.APIGetUserWorkoutsInput) (output model.APIGetUserWorkoutsOutput) {
	listInput := model.ListInput{}
	listInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "WorkoutLogs", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	workoutOutputs, _, err := r.workoutService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// output 轉換成 data item
	dataItems := make([]*model.APIGetUserWorkoutsDataItem, 0)
	for _, workoutOutput := range workoutOutputs {
		dataItem := model.APIGetUserWorkoutsDataItem{}
		if err := util.Parser(workoutOutput, &dataItem); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		dataItem.Finish = util.PointerInt(0)
		if len(*workoutOutput.WorkoutLogs) > 0 {
			dataItem.Finish = util.PointerInt(1)
		}
		dataItems = append(dataItems, &dataItem)
	}
	// parser output
	data := model.APIGetUserWorkoutsData{}
	if err := util.Parser(dataItems, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}

func (r *resolver) APIUpdateUserWorkout(tx *gorm.DB, input *model.APIUpdateUserWorkoutInput) (output model.APIUpdateUserWorkoutOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此訓練擁有者，無法修改資源")
		return output
	}
	// 查詢訓練資訊
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err := r.workoutService.Tx(tx).Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.workoutService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 上傳 start audio
	if input.Form.StartAudio != nil {
		// 儲存新 start audio
		startAudioNamed, err := r.startAudioTool.Save(input.Form.StartAudio.Data, input.Form.StartAudio.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.StartAudio = util.PointerString(startAudioNamed)
		if err := r.workoutService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊 start audio
		_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.StartAudio, ""))
	}
	// 上傳 end audio
	if input.Form.EndAudio != nil {
		// 儲存新 end audio
		endAudioNamed, err := r.endAudioTool.Save(input.Form.EndAudio.Data, input.Form.EndAudio.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.EndAudio = util.PointerString(endAudioNamed)
		if err := r.workoutService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊 end audio
		_ = r.endAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.EndAudio, ""))
	}
	tx.Commit()
	// parser output
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err = r.workoutService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIUpdateUserWorkoutData{}
	if err := util.Parser(workoutOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Data = &data
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteUserWorkoutStartAudio(input *model.APIDeleteUserWorkoutStartAudioInput) (output model.APIDeleteUserWorkoutStartAudioOutput) {
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.ID)
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
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err := r.workoutService.Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.StartAudio = util.PointerString("")
	if err := r.workoutService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除音檔
	_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.StartAudio, ""))
	// parser output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteUserWorkoutEndAudio(input *model.APIDeleteUserWorkoutEndAudioInput) (output model.APIDeleteUserWorkoutEndAudioOutput) {
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.ID)
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
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err := r.workoutService.Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.EndAudio = util.PointerString("")
	if err := r.workoutService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除音檔
	_ = r.endAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.EndAudio, ""))
	// parser output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetTrainerWorkouts(input *model.APIGetTrainerWorkoutsInput) (output model.APIGetTrainerWorkoutsOutput) {
	listInput := model.ListInput{}
	listInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutOutputs, _, err := r.workoutService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetTrainerWorkoutsData{}
	if err := util.Parser(workoutOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APICreateTrainerWorkout(tx *gorm.DB, input *model.APICreateTrainerWorkoutInput) (output model.APICreateTrainerWorkoutOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證課表權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.WorkoutCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達訓練數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法創建資源")
		return output
	}
	// 創建訓練
	workoutTable := model.Table{}
	workoutTable.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutTable.Name = util.PointerString(input.Body.Name)
	workoutTable.Equipment = util.PointerString("")
	workoutTable.StartAudio = util.PointerString("")
	workoutTable.EndAudio = util.PointerString("")
	workoutTable.WorkoutSetCount = util.PointerInt(0)
	workoutID, err := r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此計畫訓練數量
	workoutListInput := model.ListInput{}
	workoutListInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此計畫訓練數量
	planTable := planModel.Table{}
	planTable.ID = util.PointerInt64(input.Uri.PlanID)
	planTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.planService.Tx(tx).Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表訓練數量
	workoutListInput = model.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err = r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此課表訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateTrainerWorkoutData{}
	data.ID = util.PointerInt64(workoutID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APICreateTrainerWorkoutFromTemplate(tx *gorm.DB, input *model.APICreateTrainerWorkoutInput) (output model.APICreateTrainerWorkoutOutput) {
	defer tx.Rollback()
	/** 驗證權限流程 */
	// 1.驗證課表權限
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.WorkoutCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達訓練數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法創建資源")
		return output
	}
	// 2.驗證模板 workout_id 合法性
	findCourseInput = courseModel.FindInput{}
	findCourseInput.WorkoutID = input.Body.WorkoutTemplateID
	templateOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.ID, 0) != util.OnNilJustReturnInt64(templateOutput.ID, 0) {
		output.Set(code.BadRequest, "非此課表所屬模板id，無法創建資源")
		return output
	}
	/** 查詢模板資訊 */
	// 查詢模板 workout 資訊
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = input.Body.WorkoutTemplateID
	workoutOutput, err := r.workoutService.Tx(tx).Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢模板 workout_set 資訊
	workoutSetListInput := workoutSetModel.ListInput{}
	workoutSetListInput.WorkoutID = input.Body.WorkoutTemplateID
	workoutSetOutputs, _, err := r.workoutSetService.Tx(tx).List(&workoutSetListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	/** 以模板創建 workout */
	// 1. 創建 workout 到指定 plan_id
	workoutTable := model.Table{}
	workoutTable.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutTable.Name = util.PointerString(input.Body.Name)
	workoutTable.Equipment = workoutOutput.Equipment
	workoutTable.StartAudio = workoutOutput.StartAudio
	workoutTable.EndAudio = workoutOutput.EndAudio
	workoutTable.WorkoutSetCount = workoutOutput.WorkoutSetCount
	newWorkoutID, err := r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2. 創建 workout_sets
	workoutSetTables := make([]*workoutSetModel.Table, 0)
	if err := util.Parser(workoutSetOutputs, &workoutSetTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, workoutSetTable := range workoutSetTables {
		workoutSetTable.ID = nil
		workoutSetTable.WorkoutID = util.PointerInt64(newWorkoutID)
	}
	_, err = r.workoutSetService.Tx(tx).Create(workoutSetTables)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	/** 更新訓練與訓練組數 */
	// 1. 查詢此課表訓練數量
	workoutListInput := model.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2. 更新此課表訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3. 查詢此計畫訓練數量
	workoutListInput = model.ListInput{}
	workoutListInput.PlanID = util.PointerInt64(input.Uri.PlanID)
	workoutOutputs, _, err = r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 4. 更新此計畫訓練數量
	planTable := planModel.Table{}
	planTable.ID = util.PointerInt64(input.Uri.PlanID)
	planTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.planService.Tx(tx).Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 5. 更新此訓練的訓練組數量
	workoutTable = model.Table{}
	workoutTable.ID = util.PointerInt64(newWorkoutID)
	workoutTable.WorkoutSetCount = util.PointerInt(len(workoutSetOutputs))
	if err := r.workoutService.Tx(tx).Update(&workoutTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateTrainerWorkoutData{}
	data.ID = util.PointerInt64(newWorkoutID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIUpdateTrainerWorkout(tx *gorm.DB, input *model.APIUpdateTrainerWorkoutInput) (output model.APIUpdateTrainerWorkoutOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此訓練擁有者，無法修改資源")
		return output
	}
	// 查詢訓練資訊
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err := r.workoutService.Tx(tx).Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改訓練
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.workoutService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 上傳 start audio
	if input.Form.StartAudio != nil {
		// 儲存新 start audio
		startAudioNamed, err := r.startAudioTool.Save(input.Form.StartAudio.Data, input.Form.StartAudio.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.StartAudio = util.PointerString(startAudioNamed)
		if err := r.workoutService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊 start audio
		_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.StartAudio, ""))
	}
	// 上傳 end audio
	if input.Form.EndAudio != nil {
		// 儲存新 end audio
		endAudioNamed, err := r.endAudioTool.Save(input.Form.EndAudio.Data, input.Form.EndAudio.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.EndAudio = util.PointerString(endAudioNamed)
		if err := r.workoutService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊 end audio
		_ = r.endAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.EndAudio, ""))
	}
	tx.Commit()
	// parser output
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err = r.workoutService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIUpdateTrainerWorkoutData{}
	if err := util.Parser(workoutOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Data = &data
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteTrainerWorkout(tx *gorm.DB, input *model.APIDeleteTrainerWorkoutInput) (output model.APIDeleteTrainerWorkoutOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢訓練
	findWorkoutInput := model.FindInput{}
	findWorkoutInput.ID = util.PointerInt64(input.Uri.ID)
	workoutOutput, err := r.workoutService.Tx(tx).Find(&findWorkoutInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證課表刪除權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout {
		output.Set(code.BadRequest, "單一訓練課表，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	// 刪除訓練
	deleteWorkoutInput := model.DeleteInput{}
	deleteWorkoutInput.ID = input.Uri.ID
	if err := r.workoutService.Tx(tx).Delete(&deleteWorkoutInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表訓練數量
	workoutListInput := model.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此課表訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此計畫訓練數量
	workoutListInput = model.ListInput{}
	workoutListInput.PlanID = workoutOutput.PlanID
	workoutOutputs, _, err = r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此計畫訓練數量
	planTable := planModel.Table{}
	planTable.ID = workoutOutput.PlanID
	planTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.planService.Tx(tx).Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除 start audio 檔案
	_ = r.startAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.StartAudio, ""))
	// 刪除 progress audio 檔案
	_ = r.endAudioTool.Delete(util.OnNilJustReturnString(workoutOutput.EndAudio, ""))

	tx.Commit()
	output.SetStatus(code.Success)
	return output
}
