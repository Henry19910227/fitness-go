package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New(mysql.NewMockSetting())
	assert.NoError(t, err)
}
