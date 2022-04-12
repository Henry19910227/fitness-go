package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	//創建訓練組紀錄
	if err := w.workoutSetLogRepo.CreateWorkoutSetLogs(tx, workoutSetLogParams); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "workout_set_log repo", err)
	}
	//計算課表統計
	userCourseStatisticModel, err := w.workoutLogRepo.CalculateUserCourseStatistic(tx, userID, workoutID)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout_log repo", err)
	}
	//計算計畫統計
	userPlanStatisticModel, err := w.workoutLogRepo.CalculateUserPlanStatistic(tx, userID, workoutID)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout_log repo", err)
	}
	//獲取 workout set log model
	workoutSetLogModels, err := w.workoutSetLogRepo.FindWorkoutSetLogsByWorkoutLogID(tx, workoutLogID)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout_set_log repo", err)
	}
	actionIDs := make([]int64, 0)
	for _, setLogModel := range workoutSetLogModels {
		if setLogModel.WorkoutSet != nil {
			if setLogModel.WorkoutSet.Action != nil {
				actionIDs = append(actionIDs, setLogModel.WorkoutSet.Action.ID)
			}
		}
	}
	//比對最佳紀錄
	actionPRs, err := w.actionPRRepo.FindActionPRs(tx, userID, actionIDs)
	if err != nil {
		return nil, w.errHandler.Set(c, "action pr repo", err)
	}
	actionPRDict := make(map[int64]*model.ActionPR)
	for _, actionPR := range actionPRs {
		actionPRDict[actionPR.ActionID] = actionPR
	}
	//比對突破最佳紀錄組
	workoutSetLogTags := make([]*dto.WorkoutSetLogTag, 0)
	for _, setLogModel := range workoutSetLogModels {
		workoutSetLogTag := dto.NewWorkoutSetLogTag(setLogModel)
		workoutSetLogTag.NewRecord = 0
		pr, ok := actionPRDict[setLogModel.WorkoutSet.Action.ID]
		if !ok {
			workoutSetLogTag.NewRecord = 1
		} else {
			if setLogModel.Duration > pr.Duration ||
				setLogModel.Distance > pr.Distance ||
				setLogModel.Weight > pr.Weight ||
				setLogModel.Reps > pr.Reps ||
				setLogModel.Incline > pr.Incline {
				workoutSetLogTag.NewRecord = 1
			}
		}
		workoutSetLogTags = append(workoutSetLogTags, &workoutSetLogTag)
	}
	//計算最佳新的紀錄
	bestActionSetLogs, err := w.workoutSetLogRepo.CalculateBestWorkoutSetLog(tx, userID, actionIDs)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout set log repo", err)
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
	if _, err := w.planStatisticRepo.SaveUserPlanStatistic(tx, &model.SaveUserPlanStatisticParam{
		UserID:             userID,
		PlanID:             userPlanStatisticModel.PlanID,
		FinishWorkoutCount: userPlanStatisticModel.FinishWorkoutCount,
		Duration:           userPlanStatisticModel.Duration,
	}); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "plan_statistic repo", err)
	}
	actionPRParams := make([]*model.CreateActionPRParam, 0)
	for _, setLog := range bestActionSetLogs {
		pr := model.CreateActionPRParam{
			ActionID: setLog.ActionID,
			Weight:   setLog.Weight,
			Reps:     setLog.Reps,
			Distance: setLog.Distance,
			Duration: setLog.Duration,
			Incline:  setLog.Incline,
		}
		actionPRParams = append(actionPRParams, &pr)
	}
	//更新最佳紀錄
	if err := w.actionPRRepo.SaveActionPRs(tx, userID, actionPRParams); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "action personal record repo", err)
	}
	//計算教練學員數量
	studentCount, err := w.trainerStatisticRepo.CalculateTrainerStudentCount(tx, course.UserID)
	if err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "trainer statistic repo", err)
	}
	//更新教練學員數量
	if err := w.trainerStatisticRepo.SaveTrainerStatistic(tx, course.UserID, &model.SaveTrainerStatisticParam{
		StudentCount: &studentCount,
	}); err != nil {
		tx.Rollback()
		return nil, w.errHandler.Set(c, "trainer statistic repo", err)
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
