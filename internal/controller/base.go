package controller

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"net/http"

	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/gin-gonic/gin"
)

// Base ...
type Base struct {
}

// JSONValidatorErrorResponse ...
func (bc *Base) JSONValidatorErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"code": 800, "data": nil, "msg": msg})
}

// JSONErrorResponse ...
func (bc *Base) JSONErrorResponse(c *gin.Context, err errcode.Error) {
	if err.Code() == 9002 {
		c.JSON(http.StatusOK, gin.H{"code": err.Code(), "data": nil, "msg": err.Msg()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"code": err.Code(), "data": nil, "msg": err.Msg()})
}

// JSONSuccessResponse ...
func (bc *Base) JSONSuccessResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": msg})
}

func (bc *Base) JSONSuccessPagingResponse(c *gin.Context, data interface{}, paging *model.Paging, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "paging": paging, "msg": msg})
}

// JSONLoginSuccessResponse ...
func (bc *Base) JSONLoginSuccessResponse(c *gin.Context, token string, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "token": token, "data": data, "msg": msg})
}
