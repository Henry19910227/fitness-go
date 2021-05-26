package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto/coursedto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Course struct {
	Base
	courseService service.Course
}

func NewCourse(baseGroup *gin.RouterGroup, courseService service.Course, trainerMiddle gin.HandlerFunc) {

	course := &Course{courseService: courseService}

	courseGroup := baseGroup.Group("/course")
	courseGroup.Use(trainerMiddle)
	courseGroup.POST("", course.CreateCourse)
}

// CreateCourse 創建課表
// @Summary 創建課表
// @Description 創建課表
// @Tags Course
// @Accept json
// @Produce json
// @Security fitness_trainer_token
// @Param json_body body validator.CreateCourseBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=coursedto.CreateResult} "創建成功!"
// @Failure 400 {object} model.ErrorResult "創建失敗"
// @Router /course [POST]
func (cc *Course) CreateCourse(c *gin.Context) {
	var header validator.TokenHeader
	var body validator.CreateCourseBody
	if err := c.ShouldBindHeader(&header); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		cc.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := cc.courseService.CreateCourseByToken(c, header.Token, &coursedto.CreateCourseParam{
		Name: body.Name,
		Level: body.Level,
		Category: body.Category,
		CategoryOther: body.CategoryOther,
		ScheduleType: body.ScheduleType,
	})
	if err != nil {
		cc.JSONErrorResponse(c, err)
		return
	}
	cc.JSONSuccessResponse(c, result, "創建成功!")
}
