package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New(mysql.NewMockSetting()))
}
