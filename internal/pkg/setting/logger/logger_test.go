package logger

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	p, err := filepath.Abs("./config/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p)
}
