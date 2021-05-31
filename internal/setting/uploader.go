package setting

import (
	"github.com/spf13/viper"
)

type uploader struct {
	vp *viper.Viper
}

func NewUploader(viperTool *viper.Viper) Uploader {
	return &uploader{viperTool}
}

// GetUploadSavePath ...
func (u *uploader) GetUploadSavePath() string {
	return u.vp.GetString("Upload.UploadRootPath")
}

// GetUploadImageAllowExts ...
func (u *uploader) GetUploadImageAllowExts() []string {
	return u.vp.GetStringSlice("App.UploadImageAllowExt")
}

// GetUploadImageMaxSize ...
func (u *uploader) GetUploadImageMaxSize() int {
	return u.vp.GetInt("App.UploadImageMaxSize")
}
