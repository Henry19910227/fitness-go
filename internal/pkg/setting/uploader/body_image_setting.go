package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var bodyImagePath = "/body/image"

type bodyImageSetting struct {
	vp *viper.Viper
}

func NewBodyImage() Setting {
	return &bodyImageSetting{vp: vp.Shared()}
}

func (c *bodyImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *bodyImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *bodyImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + bodyImagePath
}
