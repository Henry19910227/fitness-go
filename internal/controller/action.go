package controller

import (
	"github.com/Henry19910227/fitness-go/internal/access"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Action struct {
	Base
	actionService service.Action
	actionAccess  access.Action
	trainerAccess access.Trainer
}

func NewAction(baseGroup *gin.RouterGroup,
	actionService service.Action,
	actionAccess  access.Action,
	trainerAccess access.Trainer,
	userMidd midd.User,
	courseMidd midd.Course) {
	baseGroup.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	baseGroup.StaticFS("/resource/action/video", http.Dir("./volumes/storage/action/video"))
	action := &Action{actionService: actionService, actionAccess: actionAccess, trainerAccess: trainerAccess}

	baseGroup.PATCH("/action/:action_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.UpdateAction)

	baseGroup.DELETE("/action/:action_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.DeleteAction)

	baseGroup.POST("/action/:action_id/cover",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.UploadActionCover)

	baseGroup.POST("/action/:action_id/video",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.UploadActionVideo)

	baseGroup.DELETE("/action/:action_id/video",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing}),
		courseMidd.CourseCreatorVerify(),
		courseMidd.UserRoleAccessCourseByStatusRange([]global.CourseStatus{global.Preparing, global.Reject}),
		action.DeleteActionVideo)
}

// UpdateAction 修改動作
// @Summary 修改動作
// @Description 修改動作
// @Tags Action
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
// @Router /action/{action_id} [PATCH]
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
			Data: file,
		}
	}
	//獲取動作影片
	file, fileHeader, _ = c.Request.FormFile("video")
	var video *dto.File
	if file != nil {
		video = &dto.File{
			FileNamed: fileHeader.Filename,
			Data: file,
		}
	}
	action, err := a.actionService.UpdateAction(c, uri.ActionID, &dto.UpdateActionParam{
		Name: form.Name,
		Category: form.Category,
		Body: form.Body,
		Equipment: form.Equipment,
		Intro: form.Intro,
		Cover: cover,
		Video: video,
	})
	if err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	a.JSONSuccessResponse(c, action, "update success!")
}

// DeleteAction 刪除動作
// @Summary 刪除動作
// @Description 刪除動作
// @Tags Action
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} model.SuccessResult{data=dto.ActionID} "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /action/{action_id} [DELETE]
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

// UploadActionCover 上傳動作封面照
// @Summary 上傳動作封面照
// @Description 查看封面照 : https://www.fitness-app.tk/api/v1/resource/action/cover/{圖片名}
// @Tags Action
// @Security fitness_token
// @Accept mpfd
// @Param action_id path int64 true "動作id"
// @Param cover formData file true "封面照"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.ActionCover} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /action/{action_id}/cover [POST]
func (a *Action) UploadActionCover(c *gin.Context) {
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("cover")
	if err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := a.actionService.UploadActionCover(c, uri.ActionID, fileHeader.Filename, file)
	if e != nil {
		a.JSONErrorResponse(c, e)
		return
	}
	a.JSONSuccessResponse(c, result, "success upload")
}

// UploadActionVideo 上傳動作影片
// @Summary 上傳動作影片
// @Description 查看影片 : https://www.fitness-app.tk/api/v1/resource/action/video/{影片名}
// @Tags Action
// @Security fitness_token
// @Accept mpfd
// @Param action_id path int64 true "動作id"
// @Param video formData file true "影片檔"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.ActionVideo} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /action/{action_id}/video [POST]
func (a *Action) UploadActionVideo(c *gin.Context) {
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("video")
	if err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := a.actionService.UploadActionVideo(c, uri.ActionID, fileHeader.Filename, file)
	if e != nil {
		a.JSONErrorResponse(c, e)
		return
	}
	a.JSONSuccessResponse(c, result, "success upload")
}

// DeleteActionVideo 刪除動作影片
// @Summary 刪除動作影片
// @Description 刪除動作影片
// @Tags Action
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作影片id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /action/{action_id}/video [DELETE]
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
