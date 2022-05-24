package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"time"
)

type userIncomeMonthlyStatistic struct {
	Base
	transactionRepo repository.Transaction
	statRepo        repository.UserIncomeMonthlyStatistic
	errHandler      errcode.Handler
}

func NewUserIncomeMonthlyStatistic(transactionRepo repository.Transaction, statRepo repository.UserIncomeMonthlyStatistic, errHandler errcode.Handler) UserCourseUsageMonthlyStatistic {
	return &userIncomeMonthlyStatistic{transactionRepo: transactionRepo, statRepo: statRepo, errHandler: errHandler}
}

func (u *userIncomeMonthlyStatistic) Update() {
	tx := u.transactionRepo.CreateTransaction()
	defer u.transactionRepo.FinishTransaction(tx)
	results, err := u.statRepo.CalculateUserIncomeMonthlyCount(tx, time.Now().Format("2006-01-02"))
	if err != nil {
		u.errHandler.Set(nil, "user_income_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
	if err := u.statRepo.Save(tx, results); err != nil {
		u.errHandler.Set(nil, "user_income_monthly_statistic repo", err)
		tx.Rollback()
		return
	}
}
