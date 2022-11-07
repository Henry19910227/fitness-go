package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course/api_create_favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_course"
)

type resolver struct {
	favoriteCourseService favorite_course.Service
}

func New(favoriteCourseService favorite_course.Service) Resolver {
	return &resolver{favoriteCourseService: favoriteCourseService}
}

func (r *resolver) APICreateFavoriteCourse(input *api_create_favorite_course.Input) (output api_create_favorite_course.Output) {
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.CourseID = util.PointerInt64(input.Uri.CourseID)
	if err := r.favoriteCourseService.Create(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}