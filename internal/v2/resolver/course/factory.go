package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseSvc := courseService.NewService(db)
	uploadTool := uploader.NewCourseCoverTool()
	return New(courseSvc, uploadTool)
}
