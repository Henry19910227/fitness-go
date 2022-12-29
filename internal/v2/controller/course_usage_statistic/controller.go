package course_usage_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course_usage_statistic"
)

type controller struct {
	resolver course_usage_statistic.Resolver
}

func New(resolver course_usage_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

func (c *controller) Statistic() {
	c.resolver.Statistic()
}
