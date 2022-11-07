package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course/api_create_favorite_course"
)

type Resolver interface {
	APICreateFavoriteCourse(input *api_create_favorite_course.Input) (output api_create_favorite_course.Output)
}
