package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var workoutEndAudioPath = "/workout/end_audio"

type workoutEndAudioSetting struct {
	vp *viper.Viper
}

func NewWorkoutEndAudio() Setting {
	return &workoutEndAudioSetting{vp: vp.Shared()}
}

func (c *workoutEndAudioSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadAudioAllowExt")
}

func (c *workoutEndAudioSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadAudioMaxSize")
}

func (c *workoutEndAudioSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + workoutEndAudioPath
}
