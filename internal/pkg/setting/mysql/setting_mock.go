package mysql

type settingMock struct {
}

func NewMock() Setting {
	return &settingMock{}
}

func (s *settingMock) GetUserName() string {
	return "henry"
}

func (s *settingMock) GetPassword() string {
	return "aaaa8027"
}

func (s *settingMock) GetHost() string {
	return "35.189.179.168:3306"
}

func (s *settingMock) GetDatabase() string {
	return "fitness"
}
