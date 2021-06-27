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
}

func NewAction(baseGroup *gin.RouterGroup, actionService service.Action, userMiddleware gin.HandlerFunc)  {
	baseGroup.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	action := &Action{actionService: actionService}
	actionGroup := baseGroup.Group("/action")
	actionGroup.Use(userMiddleware)
	actionGroup.PATCH("/:action_id", action.UpdateAction)
	actionGroup.DELETE("/:action_id", action.DeleteAction)
	actionGroup.POST("/:action_id/cover", action.UploadActionCover)
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
	action, err := a.actionService.UpdateActionByToken(c, header.Token, uri.ActionID, &actiondto.UpdateActionParam{
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
	result, err := a.actionService.DeleteActionByToken(c, header.Token, uri.ActionID)
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
// @Success 200 {object} model.SuccessResult{data=coursedto.CourseCover} "成功!"
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
	file, fileHeader, err := c.Request.FormFile("cover")
	if err != nil {
		a.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := a.actionService.UploadActionCoverByToken(c, header.Token, uri.ActionID, fileHeader.Filename, file)
	if e != nil {
		a.JSONErrorResponse(c, e)
		return
	}
	a.JSONSuccessResponse(c, result, "success upload")
}
