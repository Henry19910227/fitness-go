package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log/api_get_user_action_workout_set_logs"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set_log"
)

type resolver struct {
	workoutSetLogService workout_set_log.Service
}

func New(workoutSetLogService workout_set_log.Service) Resolver {
	return &resolver{workoutSetLogService: workoutSetLogService}
}

func (r *resolver) APIGetUserActionWorkoutSetLogs(input *api_get_user_action_workout_set_logs.Input) (output api_get_user_action_workout_set_logs.Output) {
	listInput := model.ListInput{}
	listInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN workout_logs ON workout_set_logs.workout_log_id = workout_logs.id"},
		{Query: "INNER JOIN workout_sets ON workout_set_logs.workout_set_id = workout_sets.id"},
	}
	listInput.Wheres = []*whereModel.Where{
		{Query: "workout_logs.user_id = ?", Args: []interface{}{input.UserID}},
		{Query: "workout_sets.action_id = ?", Args: []interface{}{input.Uri.ActionID}},
		{Query: "workout_set_logs.create_at BETWEEN ? AND ?", Args: []interface{}{input.Query.StartDate, input.Query.EndDate}},
	}
	listInput.Page = input.Query.Page
	listInput.Size = input.Query.Size
	logOutputs, page, err := r.workoutSetLogService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	data := api_get_user_action_workout_set_logs.Data{}
	if err := util.Parser(logOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	output.Paging = page
	return output
}
