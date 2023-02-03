package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	migrateTool := migrate.NewTool()
	return New(migrateTool)
}
