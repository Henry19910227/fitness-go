package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	workoutModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"gorm.io/gorm"
)

type resolver struct {
	planService    plan.Service
	courseService  course.Service
	workoutService workout.Service
}

func New(planService plan.Service, courseService course.Service, workoutService workout.Service) Resolver {
	return &resolver{planService: planService, courseService: courseService, workoutService: workoutService}
}

func (r *resolver) APIGetCMSPlans(input *model.APIGetCMSPlansInput) interface{} {
	// parser input
	listInput := model.ListInput{}
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	if err := util.Parser(input.Query, &listInput); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Workouts", OrderBy: order_by.NewInput("create_at", "DESC")},
	}
	// 調用 repo
	result, page, err := r.planService.List(&listInput)
	if err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSPlansData{}
	if err := util.Parser(result, &data); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSPlansOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}

func (r *resolver) APICreateUserPlan(tx *gorm.DB, input *model.APICreateUserPlanInput) (output model.APICreateUserPlanOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.PlanCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達計畫數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法新增資源")
		return output
	}
	// 創建計畫
	planTable := model.Table{}
	planTable.CourseID = util.PointerInt64(input.Uri.CourseID)
	planTable.Name = util.PointerString(input.Body.Name)
	planTable.WorkoutCount = util.PointerInt(0)
	planID, err := r.planService.Tx(tx).Create(&planTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢該課表下計畫個數
	planListInput := model.ListInput{}
	planListInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	planOutputs, _, err := r.planService.Tx(tx).List(&planListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新課表計畫個數
	courseTable := courseModel.Table{}
	courseTable.ID = util.PointerInt64(input.Uri.CourseID)
	courseTable.PlanCount = util.PointerInt(len(planOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateUserPlanData{}
	data.ID = util.PointerInt64(planID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetTrainerPlans(input *model.APIGetTrainerPlansInput) (output model.APIGetTrainerPlansOutput) {
	listInput := model.ListInput{}
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "UserPlanStatistic", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	planOutputs, _, err := r.planService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetTrainerPlansData{}
	if err := util.Parser(planOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIDeleteUserPlan(tx *gorm.DB, input *model.APIDeleteUserPlanInput) (output model.APIDeleteUserPlanOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證計畫刪除權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法刪除資源")
		return output
	}
	// 刪除計畫
	deletePlanInput := model.DeleteInput{}
	deletePlanInput.ID = input.Uri.ID
	if err := r.planService.Tx(tx).Delete(&deletePlanInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表計畫數量
	planListInput := model.ListInput{}
	planListInput.CourseID = courseOutput.ID
	planOutputs, _, err := r.planService.Tx(tx).List(&planListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表訓練數量
	workoutListInput := workoutModel.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此課表計畫與訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.PlanCount = util.PointerInt(len(planOutputs))
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetUserPlans(input *model.APIGetUserPlansInput) (output model.APIGetUserPlansOutput) {
	listInput := model.ListInput{}
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "UserPlanStatistic", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	planOutputs, _, err := r.planService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserPlansData{}
	if err := util.Parser(planOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}

func (r *resolver) APIUpdateUserPlan(input *model.APIUpdateUserPlanInput) (output model.APIUpdateUserPlanOutput) {
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證計畫刪除權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != courseModel.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法修改資源")
		return output
	}
	// 修改計畫
	planTable := model.Table{}
	planTable.ID = util.PointerInt64(input.Uri.ID)
	planTable.Name = util.PointerString(input.Body.Name)
	if err := r.planService.Update(&planTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APICreateTrainerPlan(tx *gorm.DB, input *model.APICreateTrainerPlanInput) (output model.APICreateTrainerPlanOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) == courseModel.SingleWorkout && util.OnNilJustReturnInt(courseOutput.PlanCount, 0) >= 1 {
		output.Set(code.BadRequest, "已達計畫數量上限，無法創建資源")
		return output
	}
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法創建資源")
		return output
	}
	// 創建計畫
	planTable := model.Table{}
	planTable.CourseID = util.PointerInt64(input.Uri.CourseID)
	planTable.Name = util.PointerString(input.Body.Name)
	planTable.WorkoutCount = util.PointerInt(0)
	planID, err := r.planService.Tx(tx).Create(&planTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢該課表下計畫個數
	planListInput := model.ListInput{}
	planListInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	planOutputs, _, err := r.planService.Tx(tx).List(&planListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新課表計畫個數
	courseTable := courseModel.Table{}
	courseTable.ID = util.PointerInt64(input.Uri.CourseID)
	courseTable.PlanCount = util.PointerInt(len(planOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateTrainerPlanData{}
	data.ID = util.PointerInt64(planID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIDeleteTrainerPlan(tx *gorm.DB, input *model.APIDeleteTrainerPlanInput) (output model.APIDeleteTrainerPlanOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.PlanID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證計畫刪除權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) == courseModel.SaleTypePersonal {
		output.Set(code.BadRequest, "無法刪除個人課表類型資源")
		return output
	}
	courseStatus := util.OnNilJustReturnInt(courseOutput.CourseStatus, 0)
	if courseStatus == courseModel.Reviewing || courseStatus == courseModel.Sale || courseStatus == courseModel.Remove {
		output.Set(code.BadRequest, "無法刪除該狀態課表資源")
		return output
	}
	// 刪除計畫
	deletePlanInput := model.DeleteInput{}
	deletePlanInput.ID = input.Uri.ID
	if err := r.planService.Tx(tx).Delete(&deletePlanInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表計畫數量
	planListInput := model.ListInput{}
	planListInput.CourseID = courseOutput.ID
	planOutputs, _, err := r.planService.Tx(tx).List(&planListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢此課表訓練數量
	workoutListInput := workoutModel.ListInput{}
	workoutListInput.CourseID = courseOutput.ID
	workoutOutputs, _, err := r.workoutService.Tx(tx).List(&workoutListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新此課表計畫與訓練數量
	courseTable := courseModel.Table{}
	courseTable.ID = courseOutput.ID
	courseTable.PlanCount = util.PointerInt(len(planOutputs))
	courseTable.WorkoutCount = util.PointerInt(len(workoutOutputs))
	if err := r.courseService.Tx(tx).Update(&courseTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetProductPlans(input *model.APIGetProductPlansInput) (output model.APIGetProductPlansOutput) {
	listInput := model.ListInput{}
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	planOutputs, _, err := r.planService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetProductPlansData{}
	if err := util.Parser(planOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
