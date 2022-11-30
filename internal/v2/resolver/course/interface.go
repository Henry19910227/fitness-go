package course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_get_trainer_course_overview"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_update_cms_courses_status"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetFavoriteCourses(input *model.APIGetFavoriteCoursesInput) (output model.APIGetFavoriteCoursesOutput)

	APIGetCMSCourses(ctx *gin.Context, input *model.APIGetCMSCoursesInput) interface{}
	APIGetCMSCourse(ctx *gin.Context, input *model.APIGetCMSCourseInput) interface{}
	APIUpdateCMSCoursesStatus(ctx *gin.Context, tx *gorm.DB, input *api_update_cms_courses_status.Input) (output api_update_cms_courses_status.Output)
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
	APIGetTrainerCourseOverview(input *api_get_trainer_course_overview.Input) (output api_get_trainer_course_overview.Output)
	APICreateTrainerCourse(input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput)
	APICreateTrainerSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput)
	APIGetTrainerCourse(input *model.APIGetTrainerCourseInput) (output model.APIGetTrainerCourseOutput)
	APIUpdateTrainerCourse(tx *gorm.DB, input *model.APIUpdateTrainerCourseInput) (output model.APIUpdateTrainerCourseOutput)
	APIDeleteTrainerCourse(input *model.APIDeleteTrainerCourseInput) (output model.APIDeleteTrainerCourseOutput)
	APISubmitTrainerCourse(input *model.APISubmitTrainerCourseInput) (output model.APISubmitTrainerCourseOutput)

	APIGetStoreCourse(input *model.APIGetStoreCourseInput) (output model.APIGetStoreCourseOutput)
	APIGetStoreCourses(input *model.APIGetStoreCoursesInput) (output model.APIGetStoreCoursesOutput)
	APIGetStoreCourseStructure(input *model.APIGetStoreCourseStructureInput) (output model.APIGetStoreCourseStructureOutput)
	APIGetStoreTrainerCourses(input *model.APIGetStoreTrainerCoursesInput) (output model.APIGetStoreTrainerCoursesOutput)
	APIGetStoreHomePage(input *model.APIGetStoreHomePageInput) (output model.APIGetStoreHomePageOutput)
}
