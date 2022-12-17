package android_version

type Output struct {
	Table
}

func (Output) TableName() string {
	return "android_versions"
}
