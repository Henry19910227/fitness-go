package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type userIncomeMonthlyStatistic struct {
	Base
	transactionRepo repository.Transaction
	statRepo   repository.UserIncomeMonthlyStatistic
	errHandler errcode.Handler
}

func NewUserIncomeMonthlyStatistic(transactionRepo repository.Transaction, statRepo repository.UserIncomeMonthlyStatistic, errHandler errcode.Handler) UserIncomeMonthlyStatistic {
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

func (u *userIncomeMonthlyStatistic) GetUserIncomeMonthlyStatistic(c *gin.Context, userID int64) (*dto.UserIncomeMonthlyStatistic, errcode.Error) {
	income := dto.UserIncomeMonthlyStatistic{}
	if err := u.statRepo.Find(userID, &income); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, u.errHandler.Set(c, "user_income_monthly_statistic repo", err)
	}
	return &income, nil
}
