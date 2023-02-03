package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/migrate/api_migrate_up_to_latest"
)

type Resolver interface {
	APIMigrateUpToLatest(input *api_migrate_up_to_latest.Input) (output api_migrate_up_to_latest.Output)
}
