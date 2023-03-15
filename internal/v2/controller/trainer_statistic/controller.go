package trainer_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/trainer_statistic"
)

type controller struct {
	resolver trainer_statistic.Resolver
}

func New(resolver trainer_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

func (c *controller) StatisticStudentCount() {
	c.resolver.StatisticStudentCount()
}
