package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var certificateImagePath = "/trainer/certificate"

type certificateImageSetting struct {
	vp *viper.Viper
}

func NewCertificateImage() Setting {
	return &certificateImageSetting{vp: vp.Shared()}
}

func (c *certificateImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *certificateImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *certificateImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + certificateImagePath
}