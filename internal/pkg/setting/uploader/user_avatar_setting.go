package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var userAvatarPath = "/user/avatar"

type userAvatarSetting struct {
	vp *viper.Viper
}

func NewUserAvatar() Setting {
	return &userAvatarSetting{vp: vp.Shared()}
}

func (c *userAvatarSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *userAvatarSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *userAvatarSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + userAvatarPath
}
