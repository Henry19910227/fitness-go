package build

import (
	"flag"
	"sync"
)

var mode string
var once sync.Once

func RunMode() string {
	once.Do(func() {
		flag.StringVar(&mode, "m", "debug", "獲取運行模式")
		flag.Parse()
	})
	return mode
}
