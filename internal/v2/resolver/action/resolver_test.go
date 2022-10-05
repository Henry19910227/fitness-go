package action

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"io/ioutil"
	"testing"
)

func TestName(t *testing.T) {
	files, _ := ioutil.ReadDir(util.RootPath() + "/volumes/storage/action/system_image/")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
