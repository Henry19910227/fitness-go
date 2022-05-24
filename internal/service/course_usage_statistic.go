package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
)

type courseUsageStatistic struct {
	Base
	transactionRepo          repository.Transaction
	courseUsageStatisticRepo repository.CourseUsageStatistic
	errHandler               errcode.Handler
}

func NewCourseUsageStatistic(transactionRepo repository.Transaction, courseUsageStatisticRepo repository.CourseUsageStatistic, errHandler errcode.Handler) CourseUsageStatistic {
	return &courseUsageStatistic{transactionRepo: transactionRepo, courseUsageStatisticRepo: courseUsageStatisticRepo, errHandler: errHandler}
}

func (cu *courseUsageStatistic) Update() {
	tx := cu.transactionRepo.CreateTransaction()
	defer cu.transactionRepo.FinishTransaction(tx)
	//計算並更新 TotalFinishWorkoutCount 欄位
	courseUsageStatisticResults, err := cu.courseUsageStatisticRepo.CalculateTotalFinishWorkoutCount(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "total_finish_workout_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 UserFinishCount 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateUserFinishCount(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "user_finish_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 MaleFinishCount 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateMaleFinishCount(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "male_finish_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 FemaleFinishCount 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateFemaleFinishCount(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "female_finish_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 FinishCountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateFinishCountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "finish_count_avg", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age13to17CountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge13to17CountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_13_17_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age18to24CountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge18to24CountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_18_24_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age25to34CountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge25to34CountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_25_34_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age35to44CountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge35to44CountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_35_44_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age45to54CountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge45to54CountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_45_54_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age55to64CountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge55to64CountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_55_64_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 Age65UpCountAvg 欄位
	courseUsageStatisticResults, err = cu.courseUsageStatisticRepo.CalculateAge65UpCountAvg(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.Save(tx, "age_65_up_count", courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
}
