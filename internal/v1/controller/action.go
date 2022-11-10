package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Action struct {
	Base
	actionService        service.Action
	workoutSetLogService service.WorkoutSetLog
}

func NewAction(baseGroup *gin.RouterGroup,
	actionService service.Action,
	workoutSetLogService service.WorkoutSetLog,
	userMidd midd.User,
	courseMidd midd.Course) {
	baseGroup.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	baseGroup.StaticFS("/resource/action/video", http.Dir("./volumes/storage/action/video"))
	action := &Action{actionService: actionService, workoutSetLogService: workoutSetLogService}

	baseGroup.PATCH("/action/:action_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.UpdateAction)

	baseGroup.GET("/action/:action_id/best_personal_record",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		action.GetActionBestPR)

	baseGroup.DELETE("/action/:action_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.DeleteAction)

	baseGroup.DELETE("/action/:action_id/video",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.DeleteActionVideo)

	baseGroup.GET("/actions",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		action.SearchActions)

	baseGroup.GET("/action/:action_id/workout_set_logs",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		action.GetWorkoutSetLogs)
}

// UpdateAction 修改動作
// @Summary 修改動作 (API已過時，更新為 /v2/trainer/action/{action_id} [PATCH])
// @Description 查看封面照 : https://www.fitopia-hub.tk/api/v1/resource/action/cover/{圖片名} 查看影片 : https://www.fitopia-hub.tk/api/v1/resource/action/video/{影片名}
// @Tags Action_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Param name formData string false "動作名稱(1~20字元)"`
// @Param category formData int false "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"`
// @Param body formData int false "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"`
// @Param equipment formData int false "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"`
// @Param intro formData string false "動作介紹(1~400字元)"`
// @Param cover formData file false "課表封面照"
// @Param video formData file false "影片檔"
// @Success 200 {object} model.SuccessResult{data=dto.Action} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /v1/action/{action_id} [PATCH]
func (a *Action) UpdateAction(c *gin.Context) {
	var uri validator.ActionIDUri
	var form validator.UpdateActionForm
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBind(&form); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//獲取動作封面
	file, fileHeader, _ := c.Request.FormFile("cover")
	var cover *dto.File
	if file != nil {
		cover = &dto.File{
			FileNamed: fileHeader.Filename,
			Data:      file,
		}
	}
	//獲取動作影片
	file, fileHeader, _ = c.Request.FormFile("video")
	var video *dto.File
	if file != nil {
		video = &dto.File{
			FileNamed: fileHeader.Filename,
			Data:      file,
		}
	}
	action, err := a.actionService.UpdateAction(c, uri.ActionID, &dto.UpdateActionParam{
		Name:      form.Name,
		Category:  form.Category,
		Body:      form.Body,
		Equipment: form.Equipment,
		Intro:     form.Intro,
		Cover:     cover,
		Video:     video,
	})
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, action, "update success!")
}

// DeleteAction 刪除動作
// @Summary 刪除動作 (API已過時，更新為 /v2/trainer/action/{action_id} [DELETE])
// @Description 刪除動作
// @Tags Action_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} model.SuccessResult{data=dto.ActionID} "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/action/{action_id} [DELETE]
func (a *Action) DeleteAction(c *gin.Context) {
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := a.actionService.DeleteAction(c, uri.ActionID)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, result, "delete success!")
}

// DeleteActionVideo 刪除動作影片
// @Summary 刪除動作影片 (API已過時，更新為 /v2/trainer/action/{action_id}/video [DELETE])
// @Description 刪除動作影片
// @Tags Action_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作影片id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /v1/action/{action_id}/video [DELETE]
func (a *Action) DeleteActionVideo(c *gin.Context) {
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := a.actionService.DeleteActionVideo(c, uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, nil, "delete success")
}

// SearchActions 搜尋動作庫的動作列表
// @Summary 搜尋動作庫的動作列表 (API已過時，更新為 /v2/user/actions [GET])
// @Description 搜尋動作庫的動作列表
// @Tags Action_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "動作名稱"
// @Param category query string false "分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)"
// @Param body query string false "身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)"
// @Param equipment query string false "器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)"
// @Success 200 {object} model.SuccessResult{data=[]dto.Action} "查詢成功!"
// @Failure 400 {object} model.ErrorResult "查詢失敗"
// @Router /v1/actions [GET]
func (a *Action) SearchActions(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var query validator.SearchActionsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var source = "1" //只選定平台動作
	actions, err := a.actionService.SearchActions(c, uid, &dto.FindActionsParam{
		Name:      query.Name,
		Source:    &source,
		Category:  query.Category,
		Body:      query.Body,
		Equipment: query.Equipment,
	})
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, actions, "success!")
}

// GetActionBestPR 獲取動作個人最佳紀錄
// @Summary 獲取動作個人最佳紀錄 (API已過時，更新為 /v2/user/action/{action_id}/best_personal_record [GET])
// @Description 獲取動作個人最佳紀錄
// @Tags Action_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} model.SuccessResult{data=dto.ActionBestPR} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/action/{action_id}/best_personal_record [GET]
func (a *Action) GetActionBestPR(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	pr, err := a.actionService.FindActionBestPR(c, uid, uri.ActionID)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, pr, "success!")
}

// GetWorkoutSetLogs 以日期獲取動作訓練組紀錄
// @Summary 以日期獲取動作訓練組紀錄 (API已過時，更新為 /v2/user/action/{action_id}/workout_set_logs [GET])
// @Description 以日期獲取動作訓練組紀錄
// @Tags Action_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Param start_date query string true "區間開始日期 YYYY-MM-DD"
// @Param end_date query string true "區間結束日期 YYYY-MM-DD"
// @Success 200 {object} model.SuccessResult{data=[]dto.WorkoutSetLogSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗"
// @Router /v1/action/{action_id}/workout_set_logs [GET]
func (a *Action) GetWorkoutSetLogs(c *gin.Context) {
	uid, e := a.GetUID(c)
	if e != nil {
		a.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var query validator.GetWorkoutSetLogsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	logs, err := a.workoutSetLogService.GetWorkoutSetLogSummaries(c, uid, uri.ActionID, query.StartDate, query.EndDate)
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, logs, "success!")
}
