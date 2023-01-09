package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type CMSTrainer struct {
	Base
	trainerService service.Trainer
	courseService  service.Course
}

func NewCMSTrainer(baseGroup *gin.RouterGroup, trainerService service.Trainer, courseService service.Course, userMiddleware middleware.User) {
	cms := &CMSTrainer{trainerService: trainerService, courseService: courseService}
	baseGroup.GET("/cms/trainers",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetTrainers)

	baseGroup.GET("/cms/trainer/:user_id",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetTrainer)

	baseGroup.PATCH("/cms/trainer/:user_id",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.UpdateTrainer)

	baseGroup.GET("/cms/trainer/:user_id/courses",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetTrainerCourses)
}

// GetTrainers 獲取教練列表
// @Summary 獲取教練列表 (API已過時，更新為 /v2/cms/trainers [GET])
// @Description 獲取教練列表
// @Tags CMS/Trainer_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id query int64 false "用戶ID"
// @Param nickname query string false "教練名稱(1~40字元)"
// @Param email query string false "教練Email"
// @Param trainer_status query string false "教練狀態(1:正常/2:審核中/3:停權/4:未啟用)"
// @Param order_field query string false "排序欄位 (create_at:創建時間)"
// @Param order_type query string false "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessPagingResult{data=[]dto.CMSTrainerSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/cms/trainers [GET]
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
// @Summary 取得教練詳細資訊 (API已過時，更新為 /v2/cms/trainer/{user_id} [GET])
// @Description 取得教練詳細資訊
// @Tags CMS/Trainer_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Success 200 {object} model.SuccessPagingResult{data=dto.CMSTrainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/cms/trainer/{user_id} [GET]
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

// UpdateTrainer 更新教練資訊
// @Summary 更新教練資訊 (API已過時，更新為 /v2/cms/trainer/{user_id} [PATCH])
// @Description 更新教練資訊
// @Tags CMS/Trainer_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Param json_body body validator.UpdateTrainerBody true "更新欄位"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/cms/trainer/{user_id} [PATCH]
func (t *CMSTrainer) UpdateTrainer(c *gin.Context) {
	t.JSONSuccessResponse(c, nil, "API版本已經過時，請立即更新!")
	return
	//var uri validator.TrainerIDUri
	//if err := c.ShouldBindUri(&uri); err != nil {
	//	t.JSONValidatorErrorResponse(c, err.Error())
	//	return
	//}
	//var body validator.UpdateTrainerBody
	//if err := c.ShouldBindJSON(&body); err != nil {
	//	t.JSONValidatorErrorResponse(c, err.Error())
	//	return
	//}
	//trainer, err := t.trainerService.UpdateTrainer(c, uri.TrainerID, &dto.UpdateTrainerParam{
	//	TrainerStatus: body.TrainerStatus,
	//	TrainerLevel:  body.TrainerLevel,
	//})
	//if err != nil {
	//	t.JSONErrorResponse(c, err)
	//	return
	//}
	//t.JSONSuccessResponse(c, trainer, "success!")
}

// GetTrainerCourses 取得教練所屬的課表
// @Summary 取得教練所屬的課表
// @Description 取得教練所屬的課表
// @Tags CMS/Trainer_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Param order_field query string false "排序欄位 (update_at:更新時間)"
// @Param order_type query string false "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=dto.CourseSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/cms/trainer/{user_id}/courses [GET]
func (t *CMSTrainer) GetTrainerCourses(c *gin.Context) {
	var uri validator.TrainerIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
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
	courses, paging, err := t.courseService.GetCourseSummariesByUID(c,
		uri.TrainerID, []int{int(global.Reviewing), int(global.Sale), int(global.Reject), int(global.Remove)}, &dto.OrderByParam{
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
	t.JSONSuccessPagingResponse(c, courses, paging, "success!")
}
