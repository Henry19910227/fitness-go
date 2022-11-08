package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course/api_create_favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course/api_delete_favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_course"
)

type resolver struct {
	favoriteCourseService favorite_course.Service
}

func New(favoriteCourseService favorite_course.Service) Resolver {
	return &resolver{favoriteCourseService: favoriteCourseService}
}

func (r *resolver) APICreateFavoriteCourse(input *api_create_favorite_course.Input) (output api_create_favorite_course.Output) {
	// 查詢課表收藏
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	favoriteOutputs, _, err := r.favoriteCourseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(favoriteOutputs) > 0 {
		output.Set(code.BadRequest, "已加入過該課表收藏")
		return output
	}
	// 新增課表收藏
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

func (r *resolver) APIDeleteFavoriteCourse(input *api_delete_favorite_course.Input) (output api_delete_favorite_course.Output) {
	// 查詢課表收藏
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	favoriteOutputs, _, err := r.favoriteCourseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(favoriteOutputs) == 0 {
		output.Set(code.BadRequest, "查無此課表收藏")
		return output
	}
	// 刪除課表收藏
	deleteInput := model.DeleteInput{}
	deleteInput.UserID = input.UserID
	deleteInput.CourseID = input.Uri.CourseID
	if err := r.favoriteCourseService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}
