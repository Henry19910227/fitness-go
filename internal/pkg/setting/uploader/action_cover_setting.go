package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var actionCoverPath = "/action/cover"

type actionCoverSetting struct {
	vp *viper.Viper
}

func NewActionCover() Setting {
	return &actionCoverSetting{vp: vp.Shared()}
}

func (c *actionCoverSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *actionCoverSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *actionCoverSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + actionCoverPath
}
