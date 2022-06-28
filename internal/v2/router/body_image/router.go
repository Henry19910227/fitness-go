package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	body "github.com/Henry19910227/fitness-go/internal/v2/controller/body_image"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := body.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/body/image", http.Dir("./volumes/storage/body/image"))
	v2.GET("/body_images", midd.Verify([]global.Role{global.UserRole}), controller.GetBodyImages)
	v2.POST("/body_image", midd.Verify([]global.Role{global.UserRole}), controller.CreateBodyImage)
}
