package banner

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	bannerService "github.com/Henry19910227/fitness-go/internal/v2/service/banner"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bannerSvc := bannerService.NewService(db)
	uploadTool := uploader.NewBannerImageTool()
	return New(bannerSvc, uploadTool)
}
