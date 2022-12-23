package android_version

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/android_version"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := android_version.NewController(orm.Shared().DB())
	v2.GET("/android_version", controller.GetAndroidVersion)
}
