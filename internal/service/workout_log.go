package service

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type workoutLog struct {
	workoutLogRepo       repository.WorkoutLog
	workoutSetLogRepo    repository.WorkoutSetLog
	workoutSetRepo       repository.WorkoutSet
	actionPRRepo         repository.ActionPR
	courseRepo           repository.Course
	courseAssetRepo      repository.UserCourseAsset
	subscribeInfoRepo    repository.UserSubscribeInfo
	courseStatisticRepo  repository.UserCourseStatistic
	planStatisticRepo    repository.UserPlanStatistic
	trainerStatisticRepo repository.TrainerStatistic
	transactionRepo      repository.Transaction
	errHandler           errcode.Handler
}

func NewWorkoutLog(workoutLogRepo repository.WorkoutLog, workoutSetLogRepo repository.WorkoutSetLog,
	workoutSetRepo repository.WorkoutSet, actionPRRepo repository.ActionPR, courseRepo repository.Course,
	courseAssetRepo repository.UserCourseAsset, subscribeInfoRepo repository.UserSubscribeInfo, courseStatisticRepo repository.UserCourseStatistic,
	planStatisticRepo repository.UserPlanStatistic, trainerStatisticRepo repository.TrainerStatistic,
	transactionRepo repository.Transaction, errHandler errcode.Handler) WorkoutLog {
	return &workoutLog{workoutLogRepo: workoutLogRepo, workoutSetLogRepo: workoutSetLogRepo,
		workoutSetRepo: workoutSetRepo, actionPRRepo: actionPRRepo, courseRepo: courseRepo,
		courseAssetRepo: courseAssetRepo, subscribeInfoRepo: subscribeInfoRepo,
		planStatisticRepo: planStatisticRepo, trainerStatisticRepo: trainerStatisticRepo,
		transactionRepo: transactionRepo, courseStatisticRepo: courseStatisticRepo, errHandler: errHandler}
}

func (w *workoutLog) CreateWorkoutLog(c *gin.Context, userID int64, workoutID int64, param *dto.WorkoutLogParam) ([]*dto.WorkoutSetLogTag, errcode.Error) {
	if param == nil {
		return nil, nil
	}
	//驗證課表是否購買或訂閱
	course := struct {
		ID       int64 `gorm:"column:id"`
		UserID   int64 `gorm:"column:user_id"`
		SaleType int   `gorm:"column:sale_type"`
	}{}
	if err := w.courseRepo.FindCourseByWorkoutID(workoutID, &course); err != nil {
		return nil, w.errHandler.Set(c, "course repo", err)
	}
	if global.SaleType(course.SaleType) == global.SaleTypeSubscribe {
		subscribeInfo, err := w.subscribeInfoRepo.FindSubscribeInfo(userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, w.errHandler.Set(c, "subscribe info repo", err)
		}
		if subscribeInfo == nil {
			return nil, w.errHandler.Custom(8999, errors.New("尚未訂閱"))
		}
		if subscribeInfo.Status == 0 {
			return nil, w.errHandler.Custom(8999, errors.New("訂閱過期"))
		}
	}
	if global.SaleType(course.SaleType) == global.SaleTypeFree || global.SaleType(course.SaleType) == global.SaleTypeCharge {
		courseAsset, err := w.courseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
			UserID:   userID,
			CourseID: course.ID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, w.errHandler.Set(c, "course asset repo", err)
		}
		if courseAsset == nil {
			return nil, w.errHandler.Custom(8999, errors.New("尚未購買此課表"))
		}
		if courseAsset.Available == 0 {
			return nil, w.errHandler.Custom(8999, errors.New("此課表無法使用"))
		}
	}
	tx := w.transactionRepo.CreateTransaction()
	var intensity int
	if param.Intensity != nil {
		intensity = *param.Intensity
	}
	var place int
	if param.Place != nil {
		place = *param.Place
	}
	//驗證是否添加此訓練範圍內的訓練組
	setIDs, err := w.workoutSetRepo.FindWorkoutSetIDsByWorkoutID(tx, workoutID)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout set repo", err)
	}
	setIDMap := make(map[int64]int64)
	for _, setID := range setIDs {
		setIDMap[setID] = setID
	}
	for _, workoutSetLogDto := range param.WorkoutSetLogs {
		//檢查加入的setID是否在此workout底下
		if _, ok := setIDMap[workoutSetLogDto.WorkoutSetID]; !ok {
			tx.Rollback()
			return nil, w.errHandler.Custom(8999, errors.New("加入了不合法的 workout set id"))
		}
	}
	//創建訓練記錄
	workoutLogID, err := w.workoutLogRepo.CreateWorkoutLog(tx, &model.CreateWorkoutLogParam{
		UserID:    userID,
		WorkoutID: workoutID,
		Duration:  param.Duration,
		Intensity: intensity,
		Place:     place,
	})
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout_log repo", err)
	}
	//創建訓練組記錄
	workoutSetLogParams := make([]*model.WorkoutSetLogParam, 0)
	for _, workoutSetLogDto := range param.WorkoutSetLogs {
		workoutSetLogParam := model.WorkoutSetLogParam{
			WorkoutLogID: workoutLogID,
			WorkoutSetID: workoutSetLogDto.WorkoutSetID,
			Weight:       workoutSetLogDto.Weight,
			Reps:         workoutSetLogDto.Reps,
			Distance:     workoutSetLogDto.Distance,
			Duration:     workoutSetLogDto.Duration,
			Incline:      workoutSetLogDto.Incline,
		}
		workoutSetLogParams = append(workoutSetLogParams, &workoutSetLogParam)
	}
	if err := w.workoutSetLogRepo.CreateWorkoutSetLogs(tx, workoutSetLogParams); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout_set_log repo", err)
	}
	//計算課表統計
	userCourseStatisticModel, err := w.workoutLogRepo.CalculateUserCourseStatistic(tx, userID, workoutID)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout_log repo", err)
	}
	if _, err := w.courseStatisticRepo.SaveUserCourseStatistic(tx, &model.SaveUserCourseStatisticParam{
		UserID:                  userID,
		CourseID:                userCourseStatisticModel.CourseID,
		FinishWorkoutCount:      userCourseStatisticModel.FinishWorkoutCount,
		TotalFinishWorkoutCount: userCourseStatisticModel.TotalFinishWorkoutCount,
		Duration:                userCourseStatisticModel.Duration,
	}); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "course_statistic repo", err)
	}
	//計算計畫統計
	userPlanStatisticModel, err := w.workoutLogRepo.CalculateUserPlanStatistic(tx, userID, workoutID)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout_log repo", err)
	}
	if _, err := w.planStatisticRepo.SaveUserPlanStatistic(tx, &model.SaveUserPlanStatisticParam{
		UserID:             userID,
		PlanID:             userPlanStatisticModel.PlanID,
		FinishWorkoutCount: userPlanStatisticModel.FinishWorkoutCount,
		Duration:           userPlanStatisticModel.Duration,
	}); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "plan_statistic repo", err)
	}
	//計算並更新教練學員數量
	studentCount, err := w.trainerStatisticRepo.CalculateTrainerStudentCount(tx, course.UserID)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "trainer statistic repo", err)
	}
	if err := w.trainerStatisticRepo.SaveTrainerStatistic(tx, course.UserID, &model.SaveTrainerStatisticParam{
		StudentCount: &studentCount,
	}); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "trainer statistic repo", err)
	}
	//將 WorkoutSetLogModel 轉換為 WorkoutSetLogTag
	workoutSetLogModels, err := w.workoutSetLogRepo.FindWorkoutSetLogsByWorkoutLogID(tx, workoutLogID)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout_set_log repo", err)
	}
	workoutSetLogTags := make([]*dto.WorkoutSetLogTag, 0)
	for _, setLogModel := range workoutSetLogModels {
		workoutSetLogTag := dto.NewWorkoutSetLogTag(setLogModel)
		workoutSetLogTags = append(workoutSetLogTags, &workoutSetLogTag)
	}
	//蒐集每個組的動作ID
	actionIDs := make([]int64, 0)
	for _, setLogModel := range workoutSetLogModels {
		if setLogModel.WorkoutSet != nil {
			if setLogModel.WorkoutSet.Action != nil {
				actionIDs = append(actionIDs, setLogModel.WorkoutSet.Action.ID)
			}
		}
	}
	//取得各個動作最佳紀錄
	actionPRs, err := w.actionPRRepo.FindActionBestPRs(tx, userID, actionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	actionPRDict := make(map[int64]*model.ActionBestPR)
	for _, actionPR := range actionPRs {
		actionPRDict[actionPR.ActionID] = actionPR
	}
	//比較最佳紀錄並打上tag(皇冠標籤)
	updateMaxRepsActionIDs := make([]int64, 0)
	updateMaxRmActionIDs := make([]int64, 0)
	updateMaxWeightActionIDs := make([]int64, 0)
	updateMinDurationActionIDs := make([]int64, 0)
	updateMaxSpeedActionIDs := make([]int64, 0)
	updateMaxDistanceActionIDs := make([]int64, 0)
	for _, setLogTag := range workoutSetLogTags {
		pr, ok := actionPRDict[setLogTag.WorkoutSet.Action.ID]
		if ok {
			//比較最多reps
			if setLogTag.Reps > pr.MaxReps {
				setLogTag.NewRecord = 1
				updateMaxRepsActionIDs = append(updateMaxRepsActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
			//比較最大rm
			rm, err := strconv.ParseFloat(fmt.Sprintf("%.1f", setLogTag.Weight*(1+0.0333*float64(setLogTag.Reps))), 64)
			if err != nil {
				tx.Rollback()
				return nil, w.errHandler.Set(c, "parser error", err)
			}
			if rm > pr.MaxRM {
				setLogTag.NewRecord = 1
				updateMaxRmActionIDs = append(updateMaxRmActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
			//比較最大重量
			if setLogTag.Weight > pr.MaxWeight {
				setLogTag.NewRecord = 1
				updateMaxWeightActionIDs = append(updateMaxWeightActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
			//比較最短時間
			if pr.MinDuration == 0 {
				setLogTag.NewRecord = 1
				updateMinDurationActionIDs = append(updateMinDurationActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
			if pr.MinDuration > 0 && setLogTag.Duration < pr.MinDuration {
				setLogTag.NewRecord = 1
				updateMinDurationActionIDs = append(updateMinDurationActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
			//比較最大速率
			speed, err := strconv.ParseFloat(fmt.Sprintf("%.1f", setLogTag.Distance*1000/float64(setLogTag.Duration)*3600/1000), 64)
			if err != nil {
				tx.Rollback()
				return nil, w.errHandler.Set(c, "parser error", err)
			}
			if speed > pr.MaxSpeed {
				setLogTag.NewRecord = 1
				updateMaxSpeedActionIDs = append(updateMaxSpeedActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
			//比較最長距離
			if setLogTag.Distance > pr.MaxDistance {
				setLogTag.NewRecord = 1
				updateMaxDistanceActionIDs = append(updateMaxDistanceActionIDs, setLogTag.WorkoutSet.Action.ID)
			}
		}
	}
	//計算最佳reps並更新
	maxRepsRecords, err := w.actionPRRepo.CalculateMaxReps(tx, userID, updateMaxRepsActionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	saveMaxRepsRecords := make([]*model.SaveMaxRepsRecord, 0)
	for _, record := range maxRepsRecords {
		param := model.SaveMaxRepsRecord{
			UserID:   userID,
			ActionID: record.ActionID,
			Reps:     record.Reps,
		}
		saveMaxRepsRecords = append(saveMaxRepsRecords, &param)
	}
	if err := w.actionPRRepo.SaveMaxRepsRecords(tx, saveMaxRepsRecords); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	//計算最佳rm並更新
	maxRmRecords, err := w.actionPRRepo.CalculateMaxRM(tx, userID, updateMaxRmActionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	saveMaxRmRecords := make([]*model.SaveMaxRmRecord, 0)
	for _, record := range maxRmRecords {
		param := model.SaveMaxRmRecord{
			UserID:   userID,
			ActionID: record.ActionID,
			RM:       record.RM,
		}
		saveMaxRmRecords = append(saveMaxRmRecords, &param)
	}
	if err := w.actionPRRepo.SaveMaxRMRecords(tx, saveMaxRmRecords); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	//計算最佳weight並更新
	maxWeightRecords, err := w.actionPRRepo.CalculateMaxWeight(tx, userID, updateMaxWeightActionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	saveMaxWeightRecords := make([]*model.SaveMaxWeightRecord, 0)
	for _, record := range maxWeightRecords {
		param := model.SaveMaxWeightRecord{
			UserID:   userID,
			ActionID: record.ActionID,
			Weight:   record.Weight,
		}
		saveMaxWeightRecords = append(saveMaxWeightRecords, &param)
	}
	if err := w.actionPRRepo.SaveMaxWeightRecords(tx, saveMaxWeightRecords); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	//計算最佳duration並更新
	minDurationRecords, err := w.actionPRRepo.CalculateMinDuration(tx, userID, updateMinDurationActionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	saveMinDurationRecords := make([]*model.SaveMinDurationRecord, 0)
	for _, record := range minDurationRecords {
		param := model.SaveMinDurationRecord{
			UserID:   userID,
			ActionID: record.ActionID,
			Duration: record.Duration,
		}
		saveMinDurationRecords = append(saveMinDurationRecords, &param)
	}
	if err := w.actionPRRepo.SaveMinDurationRecords(tx, saveMinDurationRecords); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	//計算最佳speed並更新
	maxSpeedRecords, err := w.actionPRRepo.CalculateMaxSpeed(tx, userID, updateMaxSpeedActionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	saveMaxSpeedRecords := make([]*model.SaveMaxSpeedRecord, 0)
	for _, record := range maxSpeedRecords {
		param := model.SaveMaxSpeedRecord{
			UserID:   userID,
			ActionID: record.ActionID,
			Speed:    record.Speed,
		}
		saveMaxSpeedRecords = append(saveMaxSpeedRecords, &param)
	}
	if err := w.actionPRRepo.SaveMaxSpeedRecords(tx, saveMaxSpeedRecords); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	//計算最佳distance並更新
	maxDistanceRecords, err := w.actionPRRepo.CalculateMaxDistance(tx, userID, updateMaxDistanceActionIDs)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	saveMaxDistanceRecords := make([]*model.SaveMaxDistanceRecord, 0)
	for _, record := range maxDistanceRecords {
		param := model.SaveMaxDistanceRecord{
			UserID:   userID,
			ActionID: record.ActionID,
			Distance: record.Distance,
		}
		saveMaxDistanceRecords = append(saveMaxDistanceRecords, &param)
	}
	if err := w.actionPRRepo.SaveMaxDistanceRecords(tx, saveMaxDistanceRecords); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	w.transactionRepo.FinishTransaction(tx)
	return workoutSetLogTags, nil
}

func (w *workoutLog) GetWorkoutLog(c *gin.Context, workoutLogID int64) (*dto.WorkoutLog, errcode.Error) {
	log, err := w.workoutLogRepo.FindWorkoutLog(workoutLogID)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout log repo", err)
	}
	logSets, err := w.workoutSetLogRepo.FindWorkoutSetLogsByWorkoutLogID(nil, log.ID)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout log set repo", err)
	}
	workoutLog := dto.NewWorkoutLog(log, logSets)
	return &workoutLog, nil
}

func (w *workoutLog) GetWorkoutLogSummaries(c *gin.Context, userID int64, startDate string, endDate string) ([]*dto.WorkoutLogSummary, errcode.Error) {
	datas, err := w.workoutLogRepo.FindWorkoutLogsByDate(userID, startDate, endDate)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout log repo", err)
	}
	workoutLogs := make([]*dto.WorkoutLogSummary, 0)
	for _, data := range datas {
		workoutLog := dto.NewWorkoutLogSummary(data)
		workoutLogs = append(workoutLogs, &workoutLog)
	}
	return workoutLogs, nil
}
