package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/ios_version"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := ios_version.NewController(orm.Shared().DB())
	v2.GET("/ios_version", controller.GetIOSVersion)
}
