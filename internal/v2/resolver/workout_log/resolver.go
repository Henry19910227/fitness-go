package workout_log

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	maxDistanceModel "github.com/Henry19910227/fitness-go/internal/v2/model/max_distance_record"
	maxRepsModel "github.com/Henry19910227/fitness-go/internal/v2/model/max_reps_record"
	maxRMModel "github.com/Henry19910227/fitness-go/internal/v2/model/max_rm_record"
	maxSpeedModel "github.com/Henry19910227/fitness-go/internal/v2/model/max_speed_record"
	maxWeightModel "github.com/Henry19910227/fitness-go/internal/v2/model/max_weight_record"
	minDurationModel "github.com/Henry19910227/fitness-go/internal/v2/model/min_duration_record"
	planModel "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	userCourseStatisticModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_statistic"
	userPlanStatisticModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_plan_statistic"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_log/api_delete_user_workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_log/api_get_user_workout_log"
	workoutSetModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	workoutSetLogModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_distance_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_reps_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_rm_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_speed_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/max_weight_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/min_duration_record"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_plan_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set_log"
	"gorm.io/gorm"
	"strconv"
)

type resolver struct {
	workoutLogService          workout_log.Service
	workoutSetLogService       workout_set_log.Service
	workoutSetService          workout_set.Service
	courseService              course.Service
	planService                plan.Service
	maxDistanceService         max_distance_record.Service
	maxRepsService             max_reps_record.Service
	maxRMService               max_rm_record.Service
	maxSpeedService            max_speed_record.Service
	maxWeightService           max_weight_record.Service
	minDurationService         min_duration_record.Service
	userCourseStatisticService user_course_statistic.Service
	userPlanStatisticService   user_plan_statistic.Service
}

func New(workoutLogService workout_log.Service, workoutSetLogService workout_set_log.Service,
	workoutSetService workout_set.Service, courseService course.Service, planService plan.Service,
	maxDistanceService max_distance_record.Service, maxRepsService max_reps_record.Service,
	maxRMService max_rm_record.Service, maxSpeedService max_speed_record.Service,
	maxWeightService max_weight_record.Service, minDurationService min_duration_record.Service,
	userCourseStatisticService user_course_statistic.Service, userPlanStatisticService user_plan_statistic.Service) Resolver {
	return &resolver{workoutLogService: workoutLogService, workoutSetLogService: workoutSetLogService,
		workoutSetService: workoutSetService, courseService: courseService, planService: planService,
		maxDistanceService: maxDistanceService, maxRepsService: maxRepsService,
		maxRMService: maxRMService, maxSpeedService: maxSpeedService,
		maxWeightService: maxWeightService, minDurationService: minDurationService,
		userCourseStatisticService: userCourseStatisticService, userPlanStatisticService: userPlanStatisticService}
}

func (r *resolver) APICreateUserWorkoutLog(tx *gorm.DB, input *model.APICreateUserWorkoutLogInput) (output model.APICreateUserWorkoutLogOutput) {
	defer tx.Rollback()
	// 查詢 workoutID 關聯的課表資訊
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢 workoutID 關聯的計畫資訊
	findPlanInput := planModel.FindInput{}
	findPlanInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	planOutput, err := r.planService.Tx(tx).Find(&findPlanInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	saleType := util.OnNilJustReturnInt(courseOutput.SaleType, 0)
	ownerID := util.OnNilJustReturnInt64(courseOutput.UserID, 0)
	if saleType == courseModel.SaleTypePersonal && ownerID != input.UserID {
		output.Set(code.BadRequest, "個人課表類型，非課表擁有者，無法創建資源")
		return output
	}
	// 驗證添加的訓練組ID是否合法
	findSetListInput := workoutSetModel.ListInput{}
	findSetListInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	workoutSetOutputs, _, err := r.workoutSetService.Tx(tx).List(&findSetListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	workoutSetIDMap := make(map[int64]int64)
	for _, workoutSetOutput := range workoutSetOutputs {
		setID := util.OnNilJustReturnInt64(workoutSetOutput.ID, 0)
		workoutSetIDMap[setID] = setID
	}
	for _, setLog := range input.Body.WorkoutSetLogs {
		//檢查加入的setID是否在此workout底下
		if _, ok := workoutSetIDMap[setLog.WorkoutSetID]; !ok {
			output.Set(code.BadRequest, "加入了不合法的 workout set id")
			return output
		}
	}
	// 新增 workout log
	workoutLogTable := model.Table{}
	workoutLogTable.UserID = util.PointerInt64(input.UserID)
	workoutLogTable.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	workoutLogTable.Duration = util.PointerInt(input.Body.Duration)
	workoutLogTable.Intensity = util.PointerInt(util.OnNilJustReturnInt(input.Body.Intensity, 0))
	workoutLogTable.Place = util.PointerInt(util.OnNilJustReturnInt(input.Body.Place, 0))
	workoutLogID, err := r.workoutLogService.Tx(tx).Create(&workoutLogTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 新增 workout set log
	workoutSetLogTables := make([]*workoutSetLogModel.Table, 0)
	for _, item := range input.Body.WorkoutSetLogs {
		workoutSetLogTable := workoutSetLogModel.Table{}
		workoutSetLogTable.WorkoutLogID = util.PointerInt64(workoutLogID)
		workoutSetLogTable.WorkoutSetID = util.PointerInt64(item.WorkoutSetID)
		workoutSetLogTable.Weight = util.PointerFloat64(item.Weight)
		workoutSetLogTable.Reps = util.PointerInt(item.Reps)
		workoutSetLogTable.Distance = util.PointerFloat64(item.Distance)
		workoutSetLogTable.Duration = util.PointerInt(item.Duration)
		workoutSetLogTable.Incline = util.PointerFloat64(item.Incline)
		workoutSetLogTables = append(workoutSetLogTables, &workoutSetLogTable)
	}
	_, err = r.workoutSetLogService.Tx(tx).Create(workoutSetLogTables)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//更新用戶課表統計
	userCourseStatInput := userCourseStatisticModel.Statistic{}
	userCourseStatInput.CourseID = courseOutput.ID
	userCourseStatInput.UserID = util.PointerInt64(input.UserID)
	if err := r.userCourseStatisticService.Tx(tx).Statistic(&userCourseStatInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//更新用戶計畫統計
	userPlanStatInput := userPlanStatisticModel.Statistic{}
	userPlanStatInput.PlanID = planOutput.ID
	userPlanStatInput.UserID = util.PointerInt64(input.UserID)
	if err := r.userPlanStatisticService.Tx(tx).Statistic(&userPlanStatInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//計算並更新教練學員數量
	//studentCount, err := w.trainerStatisticRepo.CalculateTrainerStudentCount(tx, course.UserID)
	// 查詢
	workoutSetLogListInput := workoutSetLogModel.ListInput{}
	workoutSetLogListInput.WorkoutLogID = util.PointerInt64(workoutLogID)
	workoutSetLogListInput.Preloads = []*preloadModel.Preload{
		{Field: "WorkoutSet"},
		{Field: "WorkoutSet.Action"},
		{Field: "WorkoutSet.Action.MaxDistanceRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "WorkoutSet.Action.MaxRepsRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "WorkoutSet.Action.MaxRMRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "WorkoutSet.Action.MaxSpeedRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "WorkoutSet.Action.MaxWeightRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "WorkoutSet.Action.MinDurationRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	workoutSetLogOutputs, _, err := r.workoutSetLogService.Tx(tx).List(&workoutSetLogListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// output 轉換成 data item
	dataItems := make([]*model.APICreateUserWorkoutLogItem, 0)
	for _, log := range workoutSetLogOutputs {
		dataItem := model.APICreateUserWorkoutLogItem{}
		dataItem.NewRecord = util.PointerInt(0)
		if err := util.Parser(log, &dataItem); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 取出參數
		currentDistance := util.OnNilJustReturnFloat64(log.Distance, 0)
		currentWeight := util.OnNilJustReturnFloat64(log.Weight, 0)
		currentReps := util.OnNilJustReturnInt(log.Reps, 0)
		currentDuration := util.OnNilJustReturnInt(log.Duration, 0)

		//驗證是否破當前 最長距離 紀錄
		maxDistance := util.OnNilJustReturnFloat64(log.WorkoutSetOnSafe().ActionOnSafe().MaxDistanceRecordOnSafe().Distance, 0)
		if currentDistance > maxDistance {
			//更新紀錄
			table := maxDistanceModel.Table{}
			table.UserID = util.PointerInt64(input.UserID)
			table.ActionID = log.WorkoutSetOnSafe().ActionID
			table.Distance = util.PointerFloat64(currentDistance)
			_, err := r.maxDistanceService.Tx(tx).CreateOrUpdate(&table)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			//標記新紀錄
			dataItem.NewRecord = util.PointerInt(1)
		}
		//驗證是否破當前 最多次數 紀錄
		maxReps := util.OnNilJustReturnInt(log.WorkoutSetOnSafe().ActionOnSafe().MaxRepsRecordOnSafe().Reps, 0)
		if currentReps > maxReps {
			//更新紀錄
			table := maxRepsModel.Table{}
			table.UserID = util.PointerInt64(input.UserID)
			table.ActionID = log.WorkoutSetOnSafe().ActionID
			table.Reps = util.PointerInt(currentReps)
			_, err := r.maxRepsService.Tx(tx).CreateOrUpdate(&table)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			//標記新紀錄
			dataItem.NewRecord = util.PointerInt(1)
		}
		//驗證是否破當前 RM 紀錄
		rm, err := strconv.ParseFloat(fmt.Sprintf("%.1f", currentWeight*(1+0.0333*float64(currentReps))), 64)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		maxRM := util.OnNilJustReturnFloat64(log.WorkoutSetOnSafe().ActionOnSafe().MaxRMRecordOnSafe().RM, 0)
		if rm > maxRM {
			//更新紀錄
			table := maxRMModel.Table{}
			table.UserID = util.PointerInt64(input.UserID)
			table.ActionID = log.WorkoutSetOnSafe().ActionID
			table.RM = util.PointerFloat64(rm)
			_, err := r.maxRMService.Tx(tx).CreateOrUpdate(&table)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			//標記新紀錄
			dataItem.NewRecord = util.PointerInt(1)
		}
		//驗證是否破當前 最大速率 紀錄
		speed, err := strconv.ParseFloat(fmt.Sprintf("%.1f", currentDistance*1000/float64(currentDuration)*3600/1000), 64)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		maxSpeed := util.OnNilJustReturnFloat64(log.WorkoutSetOnSafe().ActionOnSafe().MaxSpeedRecordOnSafe().Speed, 0)
		if speed > maxSpeed {
			//更新紀錄
			table := maxSpeedModel.Table{}
			table.UserID = util.PointerInt64(input.UserID)
			table.ActionID = log.WorkoutSetOnSafe().ActionID
			table.Speed = util.PointerFloat64(speed)
			_, err := r.maxSpeedService.Tx(tx).CreateOrUpdate(&table)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			//標記新紀錄
			dataItem.NewRecord = util.PointerInt(1)
		}
		//驗證是否破當前 最大重量 紀錄
		maxWeight := util.OnNilJustReturnFloat64(log.WorkoutSetOnSafe().ActionOnSafe().MaxWeightRecordOnSafe().Weight, 0)
		if currentWeight > maxWeight {
			//更新紀錄
			table := maxWeightModel.Table{}
			table.UserID = util.PointerInt64(input.UserID)
			table.ActionID = log.WorkoutSetOnSafe().ActionID
			table.Weight = util.PointerFloat64(currentWeight)
			_, err := r.maxWeightService.Tx(tx).CreateOrUpdate(&table)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			//標記新紀錄
			dataItem.NewRecord = util.PointerInt(1)
		}
		//驗證是否破當前 最短時長 紀錄
		minDuration := util.OnNilJustReturnInt(log.WorkoutSetOnSafe().ActionOnSafe().MinDurationRecordOnSafe().Duration, 9999999)
		if currentDuration < minDuration {
			//更新紀錄
			table := minDurationModel.Table{}
			table.UserID = util.PointerInt64(input.UserID)
			table.ActionID = log.WorkoutSetOnSafe().ActionID
			table.Duration = util.PointerInt(currentDuration)
			_, err := r.minDurationService.Tx(tx).CreateOrUpdate(&table)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			//標記新紀錄
			dataItem.NewRecord = util.PointerInt(1)
		}
		dataItems = append(dataItems, &dataItem)
	}
	tx.Commit()
	// Parser Output
	data := model.APICreateUserWorkoutLogData{}
	if err := util.Parser(dataItems, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetUserWorkoutLogs(input *model.APIGetUserWorkoutLogsInput) (output model.APIGetUserWorkoutLogsOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Workout"},
	}
	listInput.Wheres = []*whereModel.Where{
		{Query: "workout_logs.create_at BETWEEN ? AND ?", Args: []interface{}{input.Query.StartDate + " 00:00:00", input.Query.EndDate + " 23:59:59"}},
	}
	listInput.Page = input.Query.Page
	listInput.Size = input.Query.Size
	workoutLogOutputs, page, err := r.workoutLogService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	data := model.APIGetUserWorkoutLogsData{}
	if err := util.Parser(workoutLogOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	output.Paging = page
	return output
}

func (r *resolver) APIGetUserWorkoutLog(input *api_get_user_workout_log.Input) (output api_get_user_workout_log.Output) {
	// 查詢訓練記錄
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.WorkoutLogID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Workout"},
		{Field: "WorkoutSetLogs.WorkoutSet.Action"},
	}
	workoutLogOutput, err := r.workoutLogService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證查詢權限
	if util.OnNilJustReturnInt64(workoutLogOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此紀錄擁有者，無法查詢該紀錄")
		return output
	}
	// 查詢課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN plans ON courses.id = plans.course_id"},
		{Query: "INNER JOIN workouts ON plans.id = workouts.plan_id"},
	}
	findCourseInput.Wheres = []*whereModel.Where{
		{Query: "workouts.id = ?", Args: []interface{}{util.OnNilJustReturnInt64(workoutLogOutput.WorkoutID, 0)}},
	}
	courseOutput, err := r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	data := api_get_user_workout_log.Data{}
	data.Course = &api_get_user_workout_log.Course{}
	if err := util.Parser(workoutLogOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(courseOutput, data.Course); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIDeleteUserWorkoutLog(input *api_delete_user_workout_log.Input) (output api_delete_user_workout_log.Output) {
	// 查詢訓練記錄
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.WorkoutLogID)
	workoutLogOutput, err := r.workoutLogService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證刪除權限
	if util.OnNilJustReturnInt64(workoutLogOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此紀錄擁有者，無法刪除該紀錄")
		return output
	}
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.WorkoutLogID
	if err := r.workoutLogService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	output.Set(code.Success, "success")
	return output
}
