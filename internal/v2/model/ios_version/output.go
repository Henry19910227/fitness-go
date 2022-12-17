package ios_version

type Output struct {
	Table
}

func (Output) TableName() string {
	return "ios_versions"
}
