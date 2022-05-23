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

func (cu *courseUsageStatistic) UpdateCourseUsageStatistic() {
	tx := cu.transactionRepo.CreateTransaction()
	defer cu.transactionRepo.FinishTransaction(tx)
	//計算並更新 TotalFinishWorkoutCount 欄位
	courseUsageStatisticResults, err := cu.courseUsageStatisticRepo.CalculateTotalFinishWorkoutCount(tx)
	if err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := cu.courseUsageStatisticRepo.SaveTotalFinishWorkoutCount(tx, courseUsageStatisticResults); err != nil {
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
	if err := cu.courseUsageStatisticRepo.SaveUserFinishCount(tx, courseUsageStatisticResults); err != nil {
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
	if err := cu.courseUsageStatisticRepo.SaveMaleFinishCount(tx, courseUsageStatisticResults); err != nil {
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
	if err := cu.courseUsageStatisticRepo.SaveFemaleFinishCount(tx, courseUsageStatisticResults); err != nil {
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
	if err := cu.courseUsageStatisticRepo.SaveFinishCountAvg(tx, courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
		return
	}
}
