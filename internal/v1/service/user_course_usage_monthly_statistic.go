package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type userCourseUsageMonthlyStatistic struct {
	Base
	transactionRepo repository.Transaction
	statRepo   repository.UserCourseUsageMonthlyStatistic
	errHandler errcode.Handler
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

func (u *userCourseUsageMonthlyStatistic) GetUserCourseUsageMonthlyStatistic(c *gin.Context, userID int64) (*dto.UserCourseUsageMonthlyStatistic, errcode.Error) {
	statistic := dto.UserCourseUsageMonthlyStatistic{}
	if err := u.statRepo.Find(userID, &statistic); err != nil && !errors.Is(err, gorm.ErrRecordNotFound)  {
		return nil, u.errHandler.Set(c, "user_course_usage_monthly_statistic repo", err)
	}
	return &statistic, nil
}
