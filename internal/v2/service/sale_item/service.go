package sale_item

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/sale_item"
)

type service struct {
	repository sale_item.Repository
}

func New(repository sale_item.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	return output, err
}
