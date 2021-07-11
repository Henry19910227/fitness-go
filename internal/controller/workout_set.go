package controller

import (
	"github.com/Henry19910227/fitness-go/internal/access"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type workoutset struct {
	Base
	workoutSetService service.WorkoutSet
	workoutSetAccess  access.WorkoutSet
	trainerAccess access.Trainer
}

func NewWorkoutSet(baseGroup *gin.RouterGroup,
	workoutSetService service.WorkoutSet,
	workoutSetAccess access.WorkoutSet,
	trainerAccess access.Trainer,
	userMiddleware gin.HandlerFunc)  {

	set := workoutset{workoutSetService: workoutSetService,
		workoutSetAccess: workoutSetAccess,
		trainerAccess: trainerAccess}
	setGroup := baseGroup.Group("/workout_set")
	setGroup.Use(userMiddleware)
	setGroup.PATCH("/:workout_set_id", set.UpdateWorkoutSet)
}

// UpdateWorkoutSet 修改訓練組
// @Summary 修改訓練組
// @Description 修改訓練組
// @Tags WorkoutSet
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param workout_set_id path int64 true "訓練組id"
// @Param json_body body validator.UpdateWorkoutSetBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=workoutdto.WorkoutSet} "更新成功!"
// @Failure 400 {object} model.ErrorResult "更新失敗"
// @Router /workout_set/{workout_set_id} [PATCH]
func (w *workoutset) UpdateWorkoutSet(c *gin.Context) {
	var header validator.TokenHeader
	var uri validator.WorkoutSetIDUri
	var body validator.UpdateWorkoutSetBody
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
	if err := w.trainerAccess.StatusVerify(c, header.Token); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	if err := w.workoutSetAccess.UpdateVerifyByWorkoutSetID(c, header.Token, uri.WorkoutSetID); err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	set, err := w.workoutSetService.UpdateWorkoutSet(c, uri.WorkoutSetID, &workoutdto.UpdateWorkoutSetParam{
		AutoNext: body.AutoNext,
		StartAudio: body.StartAudio,
		Remark: body.Remark,
		Weight: body.Weight,
		Reps: body.Reps,
		Distance: body.Distance,
		Duration: body.Duration,
		Incline: body.Incline,
	})
	if err != nil {
		w.JSONErrorResponse(c, err)
		return
	}
	w.JSONSuccessResponse(c, set, "update success!")
}