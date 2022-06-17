package orm

import (
	"fmt"
	mysqlDB "github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type tool struct {
	db *gorm.DB
}

func New(setting mysqlDB.Setting) Tool {
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v", setting.GetUserName(), setting.GetPassword(), setting.GetHost(), setting.GetDatabase())
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &tool{db: db.Debug()}
}

func (t *tool) DB() *gorm.DB {
	return t.db
}
