package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/migrate"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := migrate.NewController(orm.Shared().DB())
	v2.PUT("/migrate/up", controller.MigrateUpToLatest)
}
