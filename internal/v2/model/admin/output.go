package admin

type Output struct {
	Table
}

func (Output) TableName() string {
	return "admins"
}
