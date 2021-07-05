package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto/actiondto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Action struct {
	Base
	actionService service.Action
	permissions service.Permissions
}

func NewAction(baseGroup *gin.RouterGroup, actionService service.Action, permissions service.Permissions, userMiddleware gin.HandlerFunc)  {
	baseGroup.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	baseGroup.StaticFS("/resource/action/video", http.Dir("./volumes/storage/action/video"))
	action := &Action{actionService: actionService, permissions: permissions}
	actionGroup := baseGroup.Group("/action")
	actionGroup.Use(userMiddleware)
	actionGroup.PATCH("/:action_id", action.UpdateAction)
	actionGroup.DELETE("/:action_id", action.DeleteAction)
	actionGroup.POST("/:action_id/cover", action.UploadActionCover)
	actionGroup.POST("/:action_id/video", action.UploadActionVideo)
}

// UpdateAction 修改動作
// @Summary 修改動作
// @Description 修改動作
// @Tags Action
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param action_id path int64 true "動作id"
// @Param json_body body validator.UpdateActionBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=actiondto.Action} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /action/{action_id} [PATCH]
func (a *Action) UpdateAction(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.ActionIDUri
	var body validator.UpdateActionBody
	if err := c.ShouldBindHeader(&header); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := a.permissions.CheckActionOwnerByActionID(c, header.Token, uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	if err := a.permissions.CourseValidationByActionID(c, "", uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	action, err := a.actionService.UpdateAction(c, uri.ActionID, &actiondto.UpdateActionParam{
		Name: body.Name,
		Category: body.Category,
		Body: body.Body,
		Equipment: body.Equipment,
		Intro: body.Intro,
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
// @Security fitness_user_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} model.SuccessResult{data=actiondto.ActionID} "刪除成功!"
// @Failure 400 {object} model.ErrorResult "刪除失敗"
// @Router /action/{action_id} [DELETE]
func (a *Action) DeleteAction(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.ActionIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := a.permissions.CheckActionOwnerByActionID(c, header.Token, uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	if err := a.permissions.CourseValidationByActionID(c, "", uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
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
// @Security fitness_user_token
// @Accept mpfd
// @Param action_id path int64 true "動作id"
// @Param cover formData file true "封面照"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=actiondto.ActionCover} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /action/{action_id}/cover [POST]
func (a *Action) UploadActionCover(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.ActionIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := a.permissions.CheckActionOwnerByActionID(c, header.Token, uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	if err := a.permissions.CourseValidationByActionID(c, "", uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
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
// @Security fitness_user_token
// @Accept mpfd
// @Param action_id path int64 true "動作id"
// @Param video formData file true "影片檔"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=actiondto.ActionVideo} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /action/{action_id}/video [POST]
func (a *Action) UploadActionVideo(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.ActionIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := a.permissions.CheckActionOwnerByActionID(c, header.Token, uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
		return
	}
	if err := a.permissions.CourseValidationByActionID(c, "", uri.ActionID); err != nil {
		a.JSONErrorResponse(c, err)
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
