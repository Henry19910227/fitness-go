package config

import (
	"path"
	"runtime"
)

func RootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}
