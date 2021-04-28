package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type migrate struct {
	migTool tool.Migrate
	migErr  errcode.Common
}

func NewMigrate(migrateTool tool.Migrate, migErr  errcode.Common) Migrate {
	return &migrate{migTool: migrateTool, migErr: migErr}
}

func (m *migrate) Version() (uint, bool, errcode.Error) {
	version, isDirty, err := m.migTool.Version()
	if err != nil {
		return 0, false, m.migErr.Custom(9999, err)
	}
	return version, isDirty, nil
}

func (m *migrate) Up() (uint, bool, errcode.Error) {
	err := m.migTool.Up(nil)
	if err != nil {
		return 0, false, m.migErr.Custom(9999, err)
	}
	return m.Version()
}

func (m *migrate) UpStep(step int) (uint, bool, errcode.Error) {
	err := m.migTool.Up(&step)
	if err != nil {
		return 0, false, m.migErr.Custom(9999, err)
	}
	return m.Version()
}

func (m *migrate) Down() errcode.Error {
	err := m.migTool.Down(nil)
	if err != nil {
		return m.migErr.Custom(9999, err)
	}
	return nil
}

func (m *migrate) DownStep(step int) errcode.Error {
	err := m.migTool.Down(&step)
	if err != nil {
		return m.migErr.Custom(9999, err)
	}
	return nil
}

func (m *migrate) Force(version int) (uint, bool, errcode.Error) {
	err := m.migTool.Force(version)
	if err != nil {
		return 0, false, m.migErr.Custom(9999, err)
	}
	return m.Version()
}

func (m *migrate) Migrate(version uint) (uint, bool, errcode.Error) {
	err := m.migTool.Migrate(version)
	if err != nil {
		return 0, false, m.migErr.Custom(9999, err)
	}
	return m.Version()
}

