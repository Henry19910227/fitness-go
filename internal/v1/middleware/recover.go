package middleware

import (
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recover(logger handler.Logger) func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		// 獲取自定義的panic
		if str, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": nil, "msg": str})
			logger.Set(c, handler.Panic, "Server", 9999, str)
			c.Abort() //終止後續調用
			return
		}
		// 獲取系統的panic
		logger.Set(c, handler.Panic, "Server", 9999, "系統發生重大錯誤!")
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": nil, "msg": "發生不知名錯誤!"})
		c.Abort() //終止後續調用
	}
}
