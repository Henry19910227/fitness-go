package mysql

type settingMock struct {
}

func NewMock() Setting {
	return &settingMock{}
}

func (s *settingMock) GetUserName() string {
	return "root"
}

func (s *settingMock) GetPassword() string {
	return "aaaa8027"
}

func (s *settingMock) GetHost() string {
	return "127.0.0.1:3306"
}

func (s *settingMock) GetDatabase() string {
	return "fitness"
}
