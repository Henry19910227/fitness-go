package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"time"
)

type userCourseUsageMonthlyStatistic struct {
	Base
	transactionRepo          repository.Transaction
	statRepo repository.UserCourseUsageMonthlyStatistic
	errHandler               errcode.Handler
}

func NewUserCourseUsageMonthlyStatistic(transactionRepo repository.Transaction, statRepo repository.UserCourseUsageMonthlyStatistic, errHandler errcode.Handler) UserCourseUsageMonthlyStatistic {
	return &userCourseUsageMonthlyStatistic{transactionRepo: transactionRepo, statRepo: statRepo, errHandler: errHandler}
}

func (u *userCourseUsageMonthlyStatistic) Update() {
	tx := u.transactionRepo.CreateTransaction()
	defer u.transactionRepo.FinishTransaction(tx)
	//計算並更新 free_usage_count 欄位
	results, err := u.statRepo.CalculateCourseUsageMonthlyCount(tx, global.SaleTypeFree, time.Now().Format("2006-01-02"))
	if err != nil {
		u.errHandler.Set(nil, "user_course_usage_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := u.statRepo.Save(tx, "free_usage_count", results); err != nil {
		u.errHandler.Set(nil, "user_course_usage_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 subscribe_usage_count 欄位
	results, err = u.statRepo.CalculateCourseUsageMonthlyCount(tx, global.SaleTypeSubscribe, time.Now().Format("2006-01-02"))
	if err != nil {
		u.errHandler.Set(nil, "user_course_usage_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := u.statRepo.Save(tx, "subscribe_usage_count", results); err != nil {
		u.errHandler.Set(nil, "user_course_usage_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
	//計算並更新 charge_usage_count 欄位
	results, err = u.statRepo.CalculateCourseUsageMonthlyCount(tx, global.SaleTypeCharge, time.Now().Format("2006-01-02"))
	if err != nil {
		u.errHandler.Set(nil, "user_course_usage_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := u.statRepo.Save(tx, "charge_usage_count", results); err != nil {
		u.errHandler.Set(nil, "user_course_usage_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
}
