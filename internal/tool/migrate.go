package tool

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migrateTool struct {
	mysqlSetting setting.Mysql
	migSetting setting.Migrate
}

func NewMigrate(mysqlSetting setting.Mysql, migSetting setting.Migrate) Migrate {
	return &migrateTool{mysqlSetting: mysqlSetting, migSetting: migSetting}
}

func (m *migrateTool) Up(step *int) error {
	mig, err := m.newMigrateItem()
	if err != nil {
		return err
	}
	defer mig.Close()
	if step == nil {
		return mig.Up()
	}
	if *step <= 0 {
		return errors.New("migrate error")
	}
	return mig.Steps(*step)
}

func (m *migrateTool) Down(step *int) error {
	mig, err := m.newMigrateItem()
	if err != nil {
		return err
	}
	defer mig.Close()
	if step == nil {
		return mig.Down()
	}
	if *step <= 0 {
		return errors.New("migrate error")
	}
	return mig.Steps(-*step)
}

func (m *migrateTool) Version() (uint, bool, error) {
	mig, err := m.newMigrateItem()
	if err != nil {
		return 0, false, err
	}
	defer mig.Close()
	return mig.Version()
}

func (m *migrateTool) Force(version int) error {
	mig, err := m.newMigrateItem()
	if err != nil {
		return err
	}
	if err := mig.Force(version); err != nil {
		return err
	}
	return nil
}

func (m *migrateTool) Migrate(version uint) error {
	mig, err := m.newMigrateItem()
	if err != nil {
		return err
	}
	if err := mig.Migrate(version); err != nil {
		return err
	}
	return nil
}

func (m *migrateTool) newMigrateItem() (*migrate.Migrate, error) {
	source := m.migSetting.DirPathSource()
	dbURL := fmt.Sprintf("mysql://%v:%v@tcp(%v)/%v", m.mysqlSetting.GetUserName(), m.mysqlSetting.GetPassword(), m.mysqlSetting.GetHost(), m.mysqlSetting.GetDatabase())

	mig, err := migrate.New(source, dbURL)
	if err != nil {
		return nil, err
	}
	return mig, nil
}