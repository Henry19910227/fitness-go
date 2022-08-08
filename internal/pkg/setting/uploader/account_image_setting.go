package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var accountImagePath = "/trainer/account_image"

type accountImageSetting struct {
	vp *viper.Viper
}

func NewAccountImage() Setting {
	return &accountImageSetting{vp: vp.Shared()}
}

func (c *accountImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *accountImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *accountImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + accountImagePath
}
