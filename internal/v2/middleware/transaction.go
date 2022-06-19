package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()
		c.Set("db", db)
		c.Set("tx", txHandle)
		c.Next()
	}
}
