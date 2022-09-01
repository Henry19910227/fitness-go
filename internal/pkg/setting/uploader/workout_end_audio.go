package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var workoutStartAudioPath = "/workout/start_audio"

type workoutStartAudioSetting struct {
	vp *viper.Viper
}

func NewWorkoutStartAudio() Setting {
	return &workoutStartAudioSetting{vp: vp.Shared()}
}

func (c *workoutStartAudioSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadAudioAllowExt")
}

func (c *workoutStartAudioSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadAudioMaxSize")
}

func (c *workoutStartAudioSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + workoutStartAudioPath
}