package user_course_usage_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Statistic(t *testing.T) {
	repo := New(orm.Shared().DB())
	err := repo.Statistic()
	assert.NoError(t, err)
}
