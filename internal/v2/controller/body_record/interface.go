package body_record

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateBodyRecord(ctx *gin.Context)
	GetBodyRecords(ctx *gin.Context)
}
