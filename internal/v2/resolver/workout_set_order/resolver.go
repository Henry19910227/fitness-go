package workout_set_order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_order"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set_order"
	"gorm.io/gorm"
)

type resolver struct {
	workoutSetOrderService workout_set_order.Service
	courseService     course.Service
}

func New(workoutSetOrderService workout_set_order.Service, courseService course.Service) Resolver {
	return &resolver{workoutSetOrderService: workoutSetOrderService, courseService: courseService}
}

func (r *resolver) APIUpdateUserWorkoutSetOrders(tx *gorm.DB, input *model.APIUpdateUserWorkoutSetOrdersInput) (output model.APIUpdateUserWorkoutSetOrdersOutput) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := courseModel.FindInput{}
	findCourseInput.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此訓練擁有者，無法修改資源")
		return output
	}
	// 刪除舊有訓練排序
	deleteInput := model.DeleteInput{}
	deleteInput.WorkoutID = input.Uri.WorkoutID
	if err := r.workoutSetOrderService.Tx(tx).Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建新的訓練排序
	tables := make([]*model.Table, 0)
	for _, item := range input.Body.WorkoutSetOrders {
		table := model.Table{}
		table.WorkoutID = util.PointerInt64(input.Uri.WorkoutID)
		table.WorkoutSetID = util.PointerInt64(item.WorkoutSetID)
		table.Seq = util.PointerInt(item.Seq)
		tables = append(tables, &table)
	}
	if err := r.workoutSetOrderService.Tx(tx).Create(tables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser Output
	output.Set(code.Success, "success")
	return output
}
