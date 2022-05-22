package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
)

type courseUsageStatistic struct {
	Base
	transactionRepo repository.Transaction
	courseUsageStatisticRepo repository.CourseUsageStatistic
	errHandler           errcode.Handler
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
	}
	if err := cu.courseUsageStatisticRepo.SaveTotalFinishWorkoutCount(tx, courseUsageStatisticResults); err != nil {
		cu.errHandler.Set(nil, "course_usage_statistic repo", err)
		tx.Rollback()
	}
}
