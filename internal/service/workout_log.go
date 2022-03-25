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
	workoutLogRepo      repository.WorkoutLog
	workoutSetLogRepo   repository.WorkoutSetLog
	workoutSetRepo      repository.WorkoutSet
	courseRepo          repository.Course
	courseAssetRepo     repository.UserCourseAsset
	subscribeInfoRepo   repository.UserSubscribeInfo
	courseStatisticRepo repository.UserCourseStatistic
	planStatisticRepo   repository.UserPlanStatistic
	transactionRepo     repository.Transaction
	errHandler          errcode.Handler
}

func NewWorkoutLog(workoutLogRepo repository.WorkoutLog, workoutSetLogRepo repository.WorkoutSetLog,
	workoutSetRepo repository.WorkoutSet, courseRepo repository.Course,
	courseAssetRepo repository.UserCourseAsset, subscribeInfoRepo repository.UserSubscribeInfo, courseStatisticRepo repository.UserCourseStatistic,
	planStatisticRepo repository.UserPlanStatistic, transactionRepo repository.Transaction, errHandler errcode.Handler) WorkoutLog {
	return &workoutLog{workoutLogRepo: workoutLogRepo, workoutSetLogRepo: workoutSetLogRepo,
		workoutSetRepo: workoutSetRepo, courseRepo: courseRepo,
		courseAssetRepo: courseAssetRepo, subscribeInfoRepo: subscribeInfoRepo, planStatisticRepo: planStatisticRepo,
		transactionRepo: transactionRepo, courseStatisticRepo: courseStatisticRepo, errHandler: errHandler}
}

func (w *workoutLog) CreateWorkoutLog(c *gin.Context, userID int64, workoutID int64, param *dto.WorkoutLogParam) errcode.Error {
	if param == nil {
		return nil
	}
	//驗證課表是否購買或訂閱
	course := struct {
		ID       int64 `gorm:"column:id"`
		UserID   int64 `gorm:"column:user_id"`
		SaleType int   `gorm:"column:sale_type"`
	}{}
	if err := w.courseRepo.FindCourseByWorkoutID(workoutID, &course); err != nil {
		return w.errHandler.Set(c, "course repo", err)
	}
	if global.SaleType(course.SaleType) == global.SaleTypeSubscribe {
		subscribeInfo, err := w.subscribeInfoRepo.FindSubscribeInfo(userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return w.errHandler.Set(c, "subscribe info repo", err)
		}
		if subscribeInfo == nil {
			return w.errHandler.Custom(8999, errors.New("尚未訂閱"))
		}
		if subscribeInfo.Status == 0 {
			return w.errHandler.Custom(8999, errors.New("訂閱過期"))
		}
	}
	if global.SaleType(course.SaleType) == global.SaleTypeFree || global.SaleType(course.SaleType) == global.SaleTypeCharge {
		courseAsset, err := w.courseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
			UserID:   userID,
			CourseID: course.ID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return w.errHandler.Set(c, "course asset repo", err)
		}
		if courseAsset == nil {
			return w.errHandler.Custom(8999, errors.New("尚未購買此課表"))
		}
		if courseAsset.Available == 0 {
			return w.errHandler.Custom(8999, errors.New("此課表無法使用"))
		}
	}
	//創建訓練記錄
	tx := w.transactionRepo.CreateTransaction()
	var intensity int
	if param.Intensity != nil {
		intensity = *param.Intensity
	}
	var place int
	if param.Place != nil {
		place = *param.Place
	}
	workoutLogID, err := w.workoutLogRepo.CreateWorkoutLog(tx, &model.CreateWorkoutLogParam{
		UserID:    userID,
		WorkoutID: workoutID,
		Duration:  param.Duration,
		Intensity: intensity,
		Place:     place,
	})
	if err != nil {
		tx.Rollback()
		return w.errHandler.Set(c, "workout_log repo", err)
	}
	//查找此訓練底下的set id
	setIDs, err := w.workoutSetRepo.FindWorkoutSetIDsByWorkoutID(workoutID)
	if err != nil {
		tx.Rollback()
		return w.errHandler.Set(c, "workout set repo", err)
	}
	setIDMap := make(map[int64]int64)
	for _, setID := range setIDs {
		setIDMap[setID] = setID
	}
	workoutSetLogs := make([]*model.WorkoutSetLog, 0)
	for _, workoutSetLogDto := range param.WorkoutSetLogs {
		//檢查加入的setID是否在此workout底下
		if _, ok := setIDMap[workoutSetLogDto.WorkoutSetID]; !ok {
			tx.Rollback()
			return w.errHandler.Custom(8999, errors.New("加入了不合法的 workout set id"))
		}
		workoutSetLog := model.WorkoutSetLog{
			WorkoutLogID: workoutLogID,
			WorkoutSetID: workoutSetLogDto.WorkoutSetID,
			Weight:       workoutSetLogDto.Weight,
			Reps:         workoutSetLogDto.Reps,
			Distance:     workoutSetLogDto.Distance,
			Duration:     workoutSetLogDto.Duration,
			Incline:      workoutSetLogDto.Incline,
		}
		workoutSetLogs = append(workoutSetLogs, &workoutSetLog)
	}
	if err := w.workoutSetLogRepo.CreateWorkoutSetLogs(tx, workoutSetLogs); err != nil {
		tx.Rollback()
		return w.errHandler.Set(c, "workout_set_log repo", err)
	}
	w.transactionRepo.FinishTransaction(tx)

	//重新刷新訓練統計
	userCourseStatisticModel, err := w.workoutLogRepo.CalculateUserCourseStatistic(userID, workoutID)
	if err != nil {
		return w.errHandler.Set(c, "workout_log repo", err)
	}
	userPlanStatisticModel, err := w.workoutLogRepo.CalculateUserPlanStatistic(userID, workoutID)
	if err != nil {
		return w.errHandler.Set(c, "workout_log repo", err)
	}
	tx = w.transactionRepo.CreateTransaction()
	if _, err := w.courseStatisticRepo.SaveUserCourseStatistic(tx, &model.SaveUserCourseStatisticParam{
		UserID:                  userID,
		CourseID:                userCourseStatisticModel.CourseID,
		FinishWorkoutCount:      userCourseStatisticModel.FinishWorkoutCount,
		TotalFinishWorkoutCount: userCourseStatisticModel.TotalFinishWorkoutCount,
		Duration:                userCourseStatisticModel.Duration,
	}); err != nil {
		tx.Rollback()
		return w.errHandler.Set(c, "course_statistic repo", err)
	}
	if _, err := w.planStatisticRepo.SaveUserPlanStatistic(tx, &model.SaveUserPlanStatisticParam{
		UserID:             userID,
		PlanID:             userPlanStatisticModel.PlanID,
		FinishWorkoutCount: userPlanStatisticModel.FinishWorkoutCount,
		Duration:           userPlanStatisticModel.Duration,
	}); err != nil {
		tx.Rollback()
		return w.errHandler.Set(c, "plan_statistic repo", err)
	}
	w.transactionRepo.FinishTransaction(tx)
	return nil
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
