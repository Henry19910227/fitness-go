package review_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/review_image"
)

type service struct {
	repository review_image.Repository
}

func New(repository review_image.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	return output, err
}
