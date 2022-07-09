package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var bannerImagePath = "/banner/image"

type bannerImageSetting struct {
	vp *viper.Viper
}

func NewBannerImage() Setting {
	return &bannerImageSetting{vp: vp.Shared()}
}

func (c *bannerImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *bannerImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *bannerImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + bannerImagePath
}
