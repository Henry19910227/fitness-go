package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var cardFrontImagePath = "/trainer/card_front_image"

type cardFrontImageSetting struct {
	vp *viper.Viper
}

func NewCardFrontImage() Setting {
	return &cardFrontImageSetting{vp: vp.Shared()}
}

func (c *cardFrontImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *cardFrontImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *cardFrontImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + cardFrontImagePath
}
