package favorite_course

import model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course"

type Service interface {
	Create(item *model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
}
