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
	APIGetUserPersonalCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput)
	APIGetUserProgressCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput)
	APIGetUserChargeCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput)
	APIDeleteUserCourse(input *model.APIDeleteUserCourseInput) (output model.APIDeleteUserCourseOutput)
	APIUpdateUserCourse(input *model.APIUpdateUserCourseInput) (output model.APIUpdateUserCourseOutput)
	APIGetUserCourse(input *model.APIGetUserCourseInput) (output model.APIGetUserCourseOutput)
	APIGetUserCourseStructure(input *model.APIGetUserCourseStructureInput) (output model.APIGetUserCourseStructureOutput)

	APIGetTrainerCourses(input *model.APIGetTrainerCoursesInput) (output model.APIGetTrainerCoursesOutput)
	APICreateTrainerCourse(input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput)
	APICreateTrainerSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput)
	APIGetTrainerCourse(input *model.APIGetTrainerCourseInput) (output model.APIGetTrainerCourseOutput)
	APIUpdateTrainerCourse(tx *gorm.DB, input *model.APIUpdateTrainerCourseInput) (output model.APIUpdateTrainerCourseOutput)
	APIDeleteTrainerCourse(input *model.APIDeleteTrainerCourseInput) (output model.APIDeleteTrainerCourseOutput)
	APISubmitTrainerCourse(input *model.APISubmitTrainerCourseInput) (output model.APISubmitTrainerCourseOutput)
}
