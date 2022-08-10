package subscribe_plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/subscribe_plan"
)

type service struct {
	repository subscribe_plan.Repository
}

func New(repository subscribe_plan.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	return output, err
}
