package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type review struct {
	courseRepo repository.Course
	errHandler errcode.Handler
}

func NewReview(courseRepo repository.Course, errHandler errcode.Handler) Review {
	return &review{courseRepo: courseRepo, errHandler: errHandler}
}

func (r *review) CourseSubmit(c *gin.Context, courseID int64) errcode.Error {
	//驗證課表填寫完整性
	entity, err := r.courseRepo.FindCourseDetailByCourseID(courseID)
	if err != nil {
		return r.errHandler.Set(c, "course repo", err)
	}
	if err := r.VerifyCourse(entity); err != nil {
		return r.errHandler.Set(c, "verify course", err)
	}
	//送審課表(測試暫時將課表狀態改為"銷售中")
	var courseStatus = 3
	if err := r.courseRepo.UpdateCourseByID(courseID, &model.UpdateCourseParam{
		CourseStatus: &courseStatus,
	}); err != nil {
		return r.errHandler.Set(c, "course repo", err)
	}
	return nil
}

func (r *review) VerifyCourse(course *model.CourseDetailEntity) error {
	if course.Sale.ID == 0 {
		return errors.New(strconv.Itoa(errcode.UpdateError))
	}
	return nil
}
