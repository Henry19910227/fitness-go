package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var feedbackImagePath = "/feedback/image"

type feedbackImageSetting struct {
	vp *viper.Viper
}

func NewFeedbackImage() Setting {
	return &feedbackImageSetting{vp: vp.Shared()}
}

func (c *feedbackImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *feedbackImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *feedbackImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + feedbackImagePath
}
