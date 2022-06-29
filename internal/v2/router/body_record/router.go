package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	body "github.com/Henry19910227/fitness-go/internal/v2/controller/body_record"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := body.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/body_record", midd.Verify([]global.Role{global.UserRole}), controller.CreateBodyRecord)
	v2.GET("/body_records", midd.Verify([]global.Role{global.UserRole}), controller.GetBodyRecords)
	v2.GET("/body_records/latest", midd.Verify([]global.Role{global.UserRole}), controller.GetBodyRecordsLatest)
	v2.PATCH("/body_record/:body_record_id", midd.Verify([]global.Role{global.UserRole}), controller.UpdateBodyRecord)
	v2.DELETE("/body_record/:body_record_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteBodyRecord)
}
