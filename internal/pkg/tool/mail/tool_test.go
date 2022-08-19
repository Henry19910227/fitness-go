package mail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTool_Send(t *testing.T) {
	tool := NewTool()
	err := tool.Send("toyokoyo199@gmail.com", "密碼重設驗證信", "Hello World")
	assert.NoError(t, err)
}
