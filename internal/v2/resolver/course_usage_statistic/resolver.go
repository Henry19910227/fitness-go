package course_usage_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_usage_statistic"
)

type resolver struct {
	statisticService course_usage_statistic.Service
}

func New(statisticService course_usage_statistic.Service) Resolver {
	return &resolver{statisticService: statisticService}
}

func (r *resolver) Statistic() {
	_ = r.statisticService.Statistic()
}
