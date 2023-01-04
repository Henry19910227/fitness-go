package admin

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/admin"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := admin.NewController(orm.Shared().DB())
	//midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/cms/login", middleware.Transaction(orm.Shared().DB()), controller.CMSLogin)
}
