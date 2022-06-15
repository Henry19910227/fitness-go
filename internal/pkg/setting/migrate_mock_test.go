package setting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockMigrate_DirPathSource(t *testing.T) {
	mig := NewMockMigrate()
	assert.Equal(t, "", mig.DirPathSource())
}
