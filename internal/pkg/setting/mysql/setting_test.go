package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	setting := New()
	assert.Equal(t, "127.0.0.1:8889", setting.GetHost())
}
