package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var cardBackImagePath = "/trainer/card_back_image"

type cardBackImageSetting struct {
	vp *viper.Viper
}

func NewCardBackImage() Setting {
	return &cardBackImageSetting{vp: vp.Shared()}
}

func (c *cardBackImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *cardBackImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *cardBackImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + cardBackImagePath
}