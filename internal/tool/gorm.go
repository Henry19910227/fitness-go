package tool

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormTool struct {
	db *gorm.DB
}

func NewGorm(setting setting.Mysql) (Gorm, error) {
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v", setting.GetUserName(), setting.GetPassword(), setting.GetHost(), setting.GetDatabase())
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &gormTool{db.Debug()}, nil
}

func NewMockGorm(conn gorm.ConnPool) (Gorm, error)  {
	db, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn: conn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &gormTool{db.Debug()}, nil
}

func (g *gormTool) DB() *gorm.DB {
	return g.db
}