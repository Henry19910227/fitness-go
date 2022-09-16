package course_release_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_release_monthly_statistic"
	"time"
)

type resolver struct {
}

func New() Resolver {
	return &resolver{}
}

func (r *resolver) APIGetCMSCourseReleaseStatistic(input *model.APIGetCMSCourseReleaseStatisticInput) (output model.APIGetCMSCourseReleaseStatisticOutput) {
	data := model.APIGetCMSCourseReleaseStatisticData{}
	data.ID = util.PointerInt64(1)
	data.Year = util.PointerInt(input.Query.Year)
	data.Month = util.PointerInt(input.Query.Month)
	data.Total = util.PointerInt(1000)
	data.Free = util.PointerInt(200)
	data.Subscribe = util.PointerInt(400)
	data.Charge = util.PointerInt(400)
	data.Aerobic = util.PointerInt(100)
	data.IntervalTraining = util.PointerInt(100)
	data.WeightTraining = util.PointerInt(200)
	data.ResistanceTraining = util.PointerInt(200)
	data.BodyweightTraining = util.PointerInt(200)
	data.OtherTraining = util.PointerInt(200)
	data.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	data.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
