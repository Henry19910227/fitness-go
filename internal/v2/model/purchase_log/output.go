package purchase_log

type Output struct {
	Table
}

func (Output) TableName() string {
	return "purchase_logs"
}
