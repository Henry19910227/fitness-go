package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var trainerAvatarPath = "/trainer/avatar"

type trainerAvatarSetting struct {
	vp *viper.Viper
}

func NewTrainerAvatar() Setting {
	return &trainerAvatarSetting{vp: vp.Shared()}
}

func (c *trainerAvatarSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *trainerAvatarSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *trainerAvatarSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + trainerAvatarPath
}
