package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Action struct {
	Base
	actionService service.Action
}

func NewAction(baseGroup *gin.RouterGroup, actionService service.Action, userMiddleware gin.HandlerFunc)  {
	action := &Action{actionService: actionService}
	actionGroup := baseGroup.Group("/action")
	actionGroup.Use(userMiddleware)
	actionGroup.DELETE("/:action_id", action.DeleteAction)
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

