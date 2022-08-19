package util

import (
	"fmt"
	"testing"
)

func TestUnixToTime(t *testing.T) {
	time := UnixToTime(1660807164)

	fmt.Println(time.Format("2006-01-02 15:04:05"))
}
