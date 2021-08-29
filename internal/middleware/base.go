package middleware

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct {
}

func (bc *Base) JSONValidatorErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": 800, "data": nil, "msg": err.Error()})
}

func (bc *Base) JSONErrorResponse(c *gin.Context, err errcode.Error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": err.Code(), "data": nil, "msg": err.Msg()})
}

func (bc *Base) JSONSuccessResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": msg})
}
