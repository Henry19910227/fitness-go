package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTool(t *testing.T) {
	tool := New()
	fmt.Println(tool.MD5Encode("test1234"))
	assert.Equal(t, "16d7a4fca7442dda3ad93c9a726597e4", tool.MD5Encode("test1234"))
}
