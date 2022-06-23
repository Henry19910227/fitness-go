package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	actionService "github.com/Henry19910227/fitness-go/internal/v2/service/action"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	actionSvc := actionService.NewService(db)
	uploadTool := uploader.NewCourseCoverTool()
	return New(actionSvc, uploadTool)
}
