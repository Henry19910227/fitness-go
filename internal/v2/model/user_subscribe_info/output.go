package user_subscribe_info

type Output struct {
	Table
}
func (Output) TableName() string {
	return "user_subscribe_infos"
}
