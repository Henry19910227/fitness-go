package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var workoutSetStartAudioPath = "/workout_set/start_audio"

type workoutSetStartAudioSetting struct {
	vp *viper.Viper
}

func NewWorkoutSetStartAudio() Setting {
	return &workoutSetStartAudioSetting{vp: vp.Shared()}
}

func (c *workoutSetStartAudioSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadAudioAllowExt")
}

func (c *workoutSetStartAudioSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadAudioMaxSize")
}

func (c *workoutSetStartAudioSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + workoutSetStartAudioPath
}
