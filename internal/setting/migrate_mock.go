package setting

import "github.com/Henry19910227/fitness-go/migrations"

type mockMigrate struct {

}

func NewMockMigrate() Migrate {
	return &mockMigrate{}
}

func (m *mockMigrate) DirPathSource() string {
	return "file://" + migrations.RootPath()
}
