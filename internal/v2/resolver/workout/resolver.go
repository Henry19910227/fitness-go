package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	planModel "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"gorm.io/gorm"
)

type resolver struct {
	workoutService workout.Service
	planService    plan.Service
	courseService  course.Service
}

func New(workoutService workout.Service, planService plan.Service, courseService course.Service) Resolver {
	return &resolver{workoutService: workoutService, planService: planService, courseService: courseService}
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
	// 驗證課表刪除權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.WorkoutCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達計畫數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.PermissionDenied, "非課表擁有者，無法創建資源")
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
		output.Set(code.PermissionDenied, "單一訓練課表，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.PermissionDenied, "非課表擁有者，無法刪除資源")
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
