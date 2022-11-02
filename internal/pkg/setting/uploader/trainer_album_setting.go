package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var trainerAlbumImagePath = "/trainer/album"

type trainerAlbumSetting struct {
	vp *viper.Viper
}

func NewTrainerAlbum() Setting {
	return &trainerAlbumSetting{vp: vp.Shared()}
}

func (c *trainerAlbumSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (c *trainerAlbumSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadImageMaxSize")
}

func (c *trainerAlbumSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + trainerAlbumImagePath
}
