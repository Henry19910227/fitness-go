package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetFavoriteCourses(input *model.APIGetFavoriteCoursesInput) (output model.APIGetFavoriteCoursesOutput)
	APIGetCMSCourses(ctx *gin.Context, input *model.APIGetCMSCoursesInput) interface{}
	APIGetCMSCourse(ctx *gin.Context, input *model.APIGetCMSCourseInput) interface{}
	APIUpdateCMSCoursesStatus(input *model.APIUpdateCMSCoursesStatusInput) (output base.Output)
	APIUpdateCMSCourseCover(input *model.APIUpdateCMSCourseCoverInput) (output model.APIUpdateCMSCourseCoverOutput)
	APICreateUserCourse(input *model.APICreateUserCourseInput) (output model.APICreateUserCourseOutput)
	APICreateUserSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateUserCourseInput) (output model.APICreateUserCourseOutput)
}
