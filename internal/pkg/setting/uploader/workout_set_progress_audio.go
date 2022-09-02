package uploader

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

var workoutSetProgressAudioPath = "/workout_set/progress_audio"

type workoutSetProgressAudioSetting struct {
	vp *viper.Viper
}

func NewWorkoutSetProgressAudio() Setting {
	return &workoutSetProgressAudioSetting{vp: vp.Shared()}
}

func (c *workoutSetProgressAudioSetting) AllowExts() []string {
	return c.vp.GetStringSlice("Upload.UploadAudioAllowExt")
}

func (c *workoutSetProgressAudioSetting) MaxSize() int {
	return c.vp.GetInt("Upload.UploadAudioMaxSize")
}

func (c *workoutSetProgressAudioSetting) FilePath() string {
	return c.vp.GetString("Resource.RootPath") + workoutSetProgressAudioPath
}
