package mysql

func NewSetting() Setting {
	return New()
}

func NewMockSetting() Setting {
	return NewMock()
}
