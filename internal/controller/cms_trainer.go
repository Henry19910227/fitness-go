package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type CMSTrainer struct {
	Base
	trainerService service.Trainer
}

func NewCMSTrainer(baseGroup *gin.RouterGroup, trainerService service.Trainer, userMiddleware middleware.User) {
	cms := &CMSTrainer{trainerService: trainerService}
	baseGroup.GET("/cms/trainers",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetTrainers)

	baseGroup.GET("/cms/trainer/:user_id",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetTrainer)
}

// GetTrainers 獲取教練列表
// @Summary 獲取教練列表
// @Description 獲取教練列表
// @Tags CMS/Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id query int64 false "用戶ID"
// @Param nickname query string false "教練名稱(1~40字元)"
// @Param email query string false "教練Email"
// @Param trainer_status query string false "教練狀態(1:正常/2:審核中/3:停權/4:未啟用)"
// @Param order_column query string false "排序欄位 (create_at:創建時間)"
// @Param order_type query string false "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=[]dto.CMSTrainerSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /cms/trainers [GET]
func (t *CMSTrainer) GetTrainers(c *gin.Context) {
	var form validator.CMSGetTrainersQuery
	if err := c.ShouldBind(&form); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var orderByQuery validator.OrderByQuery
	if err := c.ShouldBind(&orderByQuery); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var pagingQuery validator.PagingQuery
	if err := c.ShouldBind(&pagingQuery); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	users, paging, err := t.trainerService.GetCMSTrainers(c, &dto.FinsCMSTrainersParam{
		UserID:        form.UserID,
		NickName:      form.Nickname,
		Email:         form.Email,
		TrainerStatus: form.TrainerStatus,
	}, &dto.OrderByParam{
		OrderField: orderByQuery.OrderField,
		OrderType:  orderByQuery.OrderType,
	}, &dto.PagingParam{
		Page: pagingQuery.Page,
		Size: pagingQuery.Size,
	})
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessPagingResponse(c, users, paging, "success!")
}

// GetTrainer 取得教練詳細資訊
// @Summary 取得教練詳細資訊
// @Description 取得教練詳細資訊
// @Tags CMS/Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Success 200 {object} model.SuccessResult{data=dto.CMSTrainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /cms/trainer/{user_id} [GET]
func (t *CMSTrainer) GetTrainer(c *gin.Context) {
	var uri validator.TrainerIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	trainer, err := t.trainerService.GetCMSTrainer(c, uri.TrainerID)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, trainer, "success!")
}
