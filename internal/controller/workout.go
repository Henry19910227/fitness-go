package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type workout struct {
	Base
	workoutService service.Workout
}

func NewWorkout(baseGroup *gin.RouterGroup, workoutService service.Workout, userMiddleware gin.HandlerFunc) {
	workout := workout{workoutService: workoutService}
	planGroup := baseGroup.Group("/workout")
	planGroup.Use(userMiddleware)
	planGroup.PATCH("/:workout_id", workout.UpdateWorkout)
	planGroup.DELETE("/:workout_id", workout.DeleteWorkout)
}

// UpdateWorkout 修改訓練
// @Summary 修改訓練
// @Description 修改訓練
// @Tags Workout
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_id path int64 true "訓練id"
// @Param json_body body validator.UpdateWorkoutBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=workoutdto.Workout} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /workout/{workout_id} [PATCH]
func (w *workout) UpdateWorkout(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	var body validator.UpdateWorkoutBody

	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
    workout, err := w.workoutService.UpdateWorkoutByToken(c, header.Token, uri.WorkoutID, &workoutdto.UpdateWorkoutParam{
		Name: body.Name,
		Equipment: body.Equipment,
	})
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, workout, "update success!")
}

// DeleteWorkout 刪除訓練
// @Summary 刪除訓練
// @Description 刪除訓練
// @Tags Workout
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_id path int64 true "訓練id"
// @Success 200 {object} model.SuccessResult{data=workoutdto.WorkoutID} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /workout/{workout_id} [DELETE]
func (w *workout) DeleteWorkout(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutIDUri
	if err := c.ShouldBindHeader(&header); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	data, err := w.workoutService.DeleteWorkoutByToken(c, header.Token, uri.WorkoutID)
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, data, "delete success!")
}


