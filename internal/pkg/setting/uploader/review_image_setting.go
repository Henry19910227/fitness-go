package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var reviewImagePath = "/course/review"

type reviewImageSetting struct {
	vp *viper.Viper
}

func NewReviewImage() Setting {
	return &reviewImageSetting{vp: vp.Shared()}
}

func (c *reviewImageSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *reviewImageSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *reviewImageSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + reviewImagePath
}
