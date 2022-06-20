package orm

import "sync"

var t Tool
var once sync.Once

func Shared() Tool {
	once.Do(func() {
		t = NewTool()
	})
	return t
}

var mt Tool
var mockOnce sync.Once

func Mock() Tool {
	mockOnce.Do(func() {
		mt = NewMockTool()
	})
	return mt
}