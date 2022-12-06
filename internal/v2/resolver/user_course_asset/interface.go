package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset/api_create_cms_course_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset/api_delete_cms_course_user"
)

type Resolver interface {
	APICreateCMSCourseUsers(input *api_create_cms_course_users.Input) (output api_create_cms_course_users.Output)
	APIDeleteCMSCourseUser(input *api_delete_cms_course_user.Input) (output api_delete_cms_course_user.Output)
}
