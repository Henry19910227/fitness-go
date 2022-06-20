package redis

import "sync"

var t Tool
var once sync.Once

func Shared() Tool {
	once.Do(func() {
		t = NewTool()
	})
	return t
}
