package banner

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/banner"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := banner.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/banner/image", http.Dir("./volumes/storage/banner/image"))
	v2.POST("/cms/banner", midd.Verify([]global.Role{global.AdminRole}), controller.CreateCMSBanner)
	v2.GET("/cms/banners", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSBanners)
}
