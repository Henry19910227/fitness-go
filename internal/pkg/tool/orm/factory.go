package orm

import (
	mysqlDB "github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"github.com/spf13/viper"
	"log"
)

func NewTool(vp *viper.Viper) Tool {
	tool, err := New(mysqlDB.New(vp))
	if err != nil {
		log.Fatalf(err.Error())
	}
	return tool
}

func NewMockTool() Tool {
	tool, err := New(mysqlDB.NewMockSetting())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return tool
}
