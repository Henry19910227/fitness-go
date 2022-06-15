package orm

import (
	"fmt"
	mysqlDB "github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type tool struct {
	db *gorm.DB
}

func New(setting mysqlDB.Setting) (Tool, error) {
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v", setting.GetUserName(), setting.GetPassword(), setting.GetHost(), setting.GetDatabase())
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &tool{db: db.Debug()}, nil
}

func (t *tool) DB() *gorm.DB {
	return t.db
}
