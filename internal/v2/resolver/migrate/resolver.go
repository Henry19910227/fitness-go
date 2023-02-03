package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/migrate/api_migrate_up_to_latest"
)

type resolver struct {
	migrateTool migrate.Tool
}

func New(migrateTool migrate.Tool) Resolver {
	return &resolver{migrateTool: migrateTool}
}

func (r *resolver) APIMigrateUpToLatest(input *api_migrate_up_to_latest.Input) (output api_migrate_up_to_latest.Output) {
	// 升級至最新版本
	if err := r.migrateTool.Up(nil); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 獲取當前版本
	version, isDirty, err := r.migrateTool.Version()
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	data := api_migrate_up_to_latest.Data{}
	data.Version = util.PointerInt(int(version))
	data.Dirty = util.PointerBool(isDirty)
	output.Data = &data
	output.Set(code.Success, "success")
	return output
}
