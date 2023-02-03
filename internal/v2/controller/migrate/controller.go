package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/migrate/api_migrate_up_to_latest"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/migrate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver migrate.Resolver
}

func New(resolver migrate.Resolver) Controller {
	return &controller{resolver: resolver}
}

// MigrateUpToLatest 將 Schema 升至最新版本
// @Summary 將 Schema 升至最新版本
// @Description 將 Schema 升至最新版本
// @Tags Migrate_v2
// @Accept json
// @Produce json
// @Success 200 {object} api_migrate_up_to_latest.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/migrate/up [PUT]
func (c *controller) MigrateUpToLatest(ctx *gin.Context) {
	input := api_migrate_up_to_latest.Input{}
	output := c.resolver.APIMigrateUpToLatest(&input)
	ctx.JSON(http.StatusOK, output)
}
