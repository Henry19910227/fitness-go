package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course"
	"github.com/spf13/viper"
)

func NewController(gormTool tool.Gorm, vp *viper.Viper) Controller {
	resolver := course.NewResolver(gormTool, vp)
	return New(resolver)
}
