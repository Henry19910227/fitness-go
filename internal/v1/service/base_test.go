package service

import (
	"database/sql"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
)

func PrepareMigrate() error {
	migrateTool := tool.NewMigrate(setting.NewMockMysql(), setting.NewMockMigrate())
	err := migrateTool.Down(nil)
	if err != nil && err.Error() != "no change"{
		return err
	}
	err = migrateTool.Up(nil)
	if err != nil && err.Error() != "no change" {
		return err
	}
	return nil
}

func PrepareGorm() (tool.Gorm, *sql.DB, error) {
	gormTool, err := tool.NewGorm(setting.NewMockMysql())
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := gormTool.DB().DB()
	if err != nil {
		return nil, nil, err
	}
	return gormTool, sqlDB, nil
}
