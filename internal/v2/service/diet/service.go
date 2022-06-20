package diet

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	repository "github.com/Henry19910227/fitness-go/internal/v2/repository/diet"
)

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	return output, err
}