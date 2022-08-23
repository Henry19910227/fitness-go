package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		defer func() {
			txHandle.Rollback()
			//if err := recover(); err != nil {
			//	txHandle.Rollback()
			//	c.JSON(http.StatusInternalServerError, gin.H{"code": code.BadRequest, "data": nil, "msg": err.(error)})
			//	c.Abort() //終止後續調用
			//	return
			//}
		}()
		c.Set("db", db)
		c.Set("tx", txHandle)
		c.Next()
	}
}
