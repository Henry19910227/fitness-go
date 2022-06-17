package redis

import (
	setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New(setting.New()))
}
