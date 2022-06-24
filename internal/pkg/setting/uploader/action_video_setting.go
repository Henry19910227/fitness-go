package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var actionVideoPath = "/action/video"

type actionVideoSetting struct {
	vp *viper.Viper
}

func NewActionVideo() Setting {
	return &actionVideoSetting{vp: vp.Shared()}
}

func (c *actionVideoSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadVideoAllowExt")
}

func (c *actionVideoSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadVideoMaxSize")
}

func (c *actionVideoSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + actionVideoPath
}
