package setting

import (
	"github.com/spf13/viper"
)

type resource struct {
	vp *viper.Viper
}

func NewResource(viperTool *viper.Viper) Resource {
	return &resource{viperTool}
}

func (r *resource) RootPath() string {
	return r.vp.GetString("Resource.RootPath")
}
