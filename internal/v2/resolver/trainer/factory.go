package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	trainerService "github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseSvc := trainerService.NewService(db)
	uploadTool := uploader.NewTrainerAvatarTool()
	return New(courseSvc, uploadTool)
}
