package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var courseCoverPath = "/course/cover"

type courseCoverSetting struct {
	vp *viper.Viper
}

func NewCourseCover() Setting {
	return &courseCoverSetting{vp: vp.Shared()}
}

func (c *courseCoverSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *courseCoverSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *courseCoverSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + courseCoverPath
}
