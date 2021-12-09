package controller

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Base ...
type Base struct {
}

func (bc *Base) GetUID(c *gin.Context) (int64, error) {
	v, exists := c.Get("uid")
	if !exists {
		return 0, errors.New(strconv.Itoa(errcode.DataNotFound))
	}
	uid, ok := v.(int64)
	if !ok {
		return 0, errors.New(strconv.Itoa(errcode.DataNotFound))
	}
	return uid, nil
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

func (bc *Base) JSONSuccessPagingResponse(c *gin.Context, data interface{}, paging *dto.		Paging, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "paging": paging, "msg": msg})
}

// JSONLoginSuccessResponse ...
func (bc *Base) JSONLoginSuccessResponse(c *gin.Context, token string, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "token": token, "data": data, "msg": msg})
}
