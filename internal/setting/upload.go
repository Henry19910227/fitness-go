package setting

import "github.com/spf13/viper"

type upload struct {
	vp *viper.Viper
}

func NewUploadLimit(viperTool *viper.Viper) Upload {
	return &upload{vp: viperTool}
}

func (u *upload) ImageAllowExts() []string {
	return u.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (u *upload) ImageMaxSize() int {
	return u.vp.GetInt("Upload.UploadImageMaxSize")
}

func (u *upload) AudioAllowExts() []string {
	return u.vp.GetStringSlice("Upload.UploadAudioAllowExt")
}

func (u *upload) AudioMaxSize() int {
	return u.vp.GetInt("Upload.UploadAudioMaxSize")
}