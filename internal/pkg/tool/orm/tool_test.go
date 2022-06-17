package orm

import (
	mysqlDB "github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tool := New(mysqlDB.NewMockSetting())
	assert.NotEmpty(t, tool)
}
