package setting

import "github.com/spf13/viper"

type uploadLimit struct {
	vp *viper.Viper
}

func NewUploadLimit(viperTool *viper.Viper) UploadLimit {
	return &uploadLimit{vp: viperTool}
}

func (u *uploadLimit) ImageAllowExts() []string {
	return u.vp.GetStringSlice("Upload.UploadImageAllowExt")
}

func (u *uploadLimit) ImageMaxSize() int {
	return u.vp.GetInt("Upload.UploadImageMaxSize")
}