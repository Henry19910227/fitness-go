package setting

type migrate struct {
	rootPath string
}

func NewMigrate(rootPath string) Migrate {
	return &migrate{rootPath: rootPath}
}

func (m *migrate) DirPathSource() string {
	return "file://" + m.rootPath + "/migrations"
}
