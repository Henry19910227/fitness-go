package trainer_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_statistic"
)

type resolver struct {
	trainerStatisticService trainer_statistic.Service
}

func New(trainerStatisticService trainer_statistic.Service) Resolver {
	return &resolver{trainerStatisticService: trainerStatisticService}
}

func (r *resolver) StatisticStudentCount() {
	_ = r.trainerStatisticService.StatisticStudentCount()
}
