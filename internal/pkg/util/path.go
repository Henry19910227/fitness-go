package util

import (
	"path"
	"runtime"
	"strings"
)

func RootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	result := strings.SplitAfter(path.Dir(filename), "fitness-go/")
	if len(result) == 0 {
		return ""
	}
	return result[0]
}
