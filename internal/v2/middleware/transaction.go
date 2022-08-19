package middleware

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		defer func() {
			if err := recover(); err != nil {
				txHandle.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"code": code.BadRequest, "data": nil, "msg": err.(error)})
				c.Abort() //終止後續調用
				return
			}
		}()
		c.Set("db", db)
		c.Set("tx", txHandle)
		c.Next()
	}
}
