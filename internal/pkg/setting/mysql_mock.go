package setting

type mockMysql struct {
}

func NewMockMysql() Mysql {
	return &mockMysql{}
}

func (setting *mockMysql) GetUserName() string {
	return "root"
}

func (setting *mockMysql) GetPassword() string {
	return "aaaa8027"
}

func (setting *mockMysql) GetHost() string {
	return "127.0.0.1:3306"
}

func (setting *mockMysql) GetDatabase() string {
	return "fitness"
}
