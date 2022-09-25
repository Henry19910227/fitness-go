package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type GenerateInput struct {
	DataAmount int
	UserID     []*base.GenerateSetting
}

type FindInput struct {
	optional.IDField
	PlanID       *int64 `json:"plan_id,omitempty"`        // 計畫 id
	WorkoutID    *int64 `json:"workout_id,omitempty"`     // 訓練 id
	WorkoutSetID *int64 `json:"workout_set_id,omitempty"` // 訓練組 id
	PreloadInput
}

type DeleteInput struct {
	required.IDField
}

type ListInput struct {
	optional.IDField
	optional.UserIDField
	optional.NameField
	optional.CourseStatusField
	optional.SaleTypeField
	optional.ScheduleTypeField
	SaleTypes           []int // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)
	IgnoredCourseStatus []int // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	IDs                 []int64
	PagingInput
	PreloadInput
	OrderByInput
}

type FavoriteListInput struct {
	optional.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

type ProgressListInput struct {
	required.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

type ChargeListInput struct {
	required.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetFavoriteCoursesInput /v2/favorite/courses [GET]
type APIGetFavoriteCoursesInput struct {
	required.UserIDField
	Form APIGetFavoriteCoursesForm
}
type APIGetFavoriteCoursesForm struct {
	PagingInput
}
