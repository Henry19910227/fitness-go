package migrate

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"github.com/Henry19910227/fitness-go/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type tool struct {
	mig *migrate.Migrate
}

func New(setting mysql.Setting) (Tool, error) {
	source := "file://" + migrations.RootPath()
	dbURL := fmt.Sprintf("mysql://%v:%v@tcp(%v)/%v", setting.GetUserName(), setting.GetPassword(), setting.GetHost(), setting.GetDatabase())
	mig, err := migrate.New(source, dbURL)
	if err != nil {
		return nil, err
	}
	return &tool{mig: mig}, nil
}

func (t *tool) Up(step *int) error {
	if step == nil {
		return t.mig.Up()
	}
	if *step <= 0 {
		return errors.New("migrate error")
	}
	return t.mig.Steps(*step)
}

func (t *tool) Down(step *int) error {

	if step == nil {
		return t.mig.Down()
	}
	if *step <= 0 {
		return errors.New("migrate error")
	}
	return t.mig.Steps(-*step)
}

func (t *tool) Version() (uint, bool, error) {
	return t.mig.Version()
}

func (t *tool) Force(version int) error {
	if err := t.mig.Force(version); err != nil {
		return err
	}
	return nil
}

func (t *tool) Migrate(version uint) error {
	if err := t.mig.Migrate(version); err != nil {
		return err
	}
	return nil
}
