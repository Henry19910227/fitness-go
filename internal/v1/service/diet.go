package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type diet struct {
	Base
	dietRepo   repository.Diet
	rdaRepo    repository.RDA
	errHandler errcode.Handler
}

func NewDiet(dietRepo repository.Diet, rdaRepo repository.RDA, errHandler errcode.Handler) Diet {
	return &diet{dietRepo: dietRepo, rdaRepo: rdaRepo, errHandler: errHandler}
}

func (d *diet) CreateDiet(c *gin.Context, userID int64, scheduleAt string) (*dto.Diet, errcode.Error) {
	//查找當前最新rda
	var rda model.RDA
	if err := d.rdaRepo.FindRDA(nil, &model.FindRDAParam{
		UserID: util.PointerInt64(userID),
	}, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &rda); err != nil {
		return nil, d.errHandler.Custom(9002, errors.New("尚未設定rda"))
	}
	//創建diet
	dietID, err := d.dietRepo.CreateDiet(nil, userID, rda.ID, scheduleAt)
	if err != nil {
		return nil, d.errHandler.Set(c, "diet repo", err)
	}
	//查找diet
	preloads := make([]*model.Preload, 0)
	preloads = append(preloads, &model.Preload{Field: "RDA"})
	data, err := d.dietRepo.FindDiet(nil, &model.FindDietParam{
		ID: util.PointerInt64(dietID),
	}, preloads)
	if err != nil {
		return nil, d.errHandler.Set(c, "diet repo", err)
	}
	//parser diet
	var diet dto.Diet
	if err := util.Parser(data, &diet); err != nil {
		return nil, d.errHandler.Set(c, "parser error", err)
	}
	return &diet, nil
}

func (d *diet) GetDiet(c *gin.Context, userID int64, scheduleAt string) (*dto.Diet, errcode.Error) {
	//查找diet
	preloads := make([]*model.Preload, 0)
	preloads = append(preloads,
		&model.Preload{Field: "RDA"},
		&model.Preload{Field: "Meals"},
		&model.Preload{Field: "Meals.Food"},
		&model.Preload{Field: "Meals.Food.FoodCategory"})
	data, err := d.dietRepo.FindDiet(nil, &model.FindDietParam{
		UserID:     util.PointerInt64(userID),
		ScheduleAt: util.PointerString(scheduleAt),
	}, preloads)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, d.errHandler.Set(c, "diet repo", err)
	}
	diet := dto.Diet{}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		diet.RDA = d.GetLatestRDA(userID)
		return &diet, nil
	}
	if err := util.Parser(data, &diet); err != nil {
		return nil, d.errHandler.Set(c, "parser error", err)
	}
	return &diet, nil
}

func (d *diet) GetLatestRDA(userID int64) *dto.RDA {
	//查找當前最新rda
	var rdaModel model.RDA
	_ = d.rdaRepo.FindRDA(nil, &model.FindRDAParam{
		UserID: util.PointerInt64(userID),
	}, &model.OrderBy{
		Field:     "create_at",
		OrderType: global.DESC,
	}, &rdaModel)
	var rda dto.RDA
	_ = util.Parser(rdaModel, &rda)
	return &rda
}
