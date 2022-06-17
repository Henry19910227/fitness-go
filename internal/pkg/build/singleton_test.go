package build

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunMode(t *testing.T) {
	assert.Equal(t, "debug", RunMode())
}
