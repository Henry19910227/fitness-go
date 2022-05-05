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
func (u *CMSTrainer) GetTrainers(c *gin.Context) {
	var form validator.CMSGetTrainersQuery
	if err := c.ShouldBind(&form); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var orderByQuery validator.OrderByQuery
	if err := c.ShouldBind(&orderByQuery); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var pagingQuery validator.PagingQuery
	if err := c.ShouldBind(&pagingQuery); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	users, paging, err := u.trainerService.GetCMSTrainers(c, &dto.FinsCMSTrainersParam{
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
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessPagingResponse(c, users, paging, "success!")
}
