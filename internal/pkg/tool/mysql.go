package tool

import (
	"database/sql"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
)

type mysqlTool struct {
	db *sql.DB
}

func NewMysql(setting setting.Mysql) (Mysql, error) {
	datasource := fmt.Sprintf("%v:%v@tcp(%v)/%v", setting.GetUserName(), setting.GetPassword(), setting.GetHost(), setting.GetDatabase())
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		return nil, err
	}
	return &mysqlTool{db}, nil
}

func (tool *mysqlTool) DB() *sql.DB {
	return tool.db
}
