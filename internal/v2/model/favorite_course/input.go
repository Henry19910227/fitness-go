package favorite_course

import "github.com/Henry19910227/fitness-go/internal/v2/field/favorite_course/required"

type DeleteInput struct {
	required.UserIDField
	required.CourseIDField
}
