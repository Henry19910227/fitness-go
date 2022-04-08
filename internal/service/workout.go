package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type workout struct {
	courseRepo     repository.Course
	workoutRepo    repository.Workout
	workoutSetRepo repository.WorkoutSet
	workoutLogRepo repository.WorkoutLog
	resHandler     handler.Resource
	uploader       handler.Uploader
	logger         handler.Logger
	jwtTool        tool.JWT
	errHandler     errcode.Handler
}

func NewWorkout(courseRepo repository.Course, workoutRepo repository.Workout, workoutSetRepo repository.WorkoutSet, workoutLogRepo repository.WorkoutLog, resHandler handler.Resource, uploader handler.Uploader, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Workout {
	return &workout{courseRepo: courseRepo, workoutRepo: workoutRepo, workoutSetRepo: workoutSetRepo, workoutLogRepo: workoutLogRepo, resHandler: resHandler, uploader: uploader, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (w *workout) CreateWorkout(c *gin.Context, planID int64, name string) (*dto.Workout, errcode.Error) {
	workoutID, err := w.workoutRepo.CreateWorkout(planID, name)
	if err != nil {
		w.logger.Set(c, handler.Error, "CourseRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	var workout dto.Workout
	if err := w.workoutRepo.FindWorkoutByID(workoutID, &workout); err != nil {
		w.logger.Set(c, handler.Error, "CourseRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &workout, nil
}

func (w *workout) GetWorkouts(c *gin.Context, planID int64) ([]*dto.Workout, errcode.Error) {
	datas, err := w.workoutRepo.FindWorkoutsByPlanID(planID)
	if err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	workouts := make([]*dto.Workout, 0)
	for _, data := range datas {
		workout := dto.Workout{
			ID:              data.ID,
			Name:            data.Name,
			Equipment:       data.Equipment,
			StartAudio:      data.StartAudio,
			EndAudio:        data.EndAudio,
			WorkoutSetCount: data.WorkoutSetCount,
		}
		workouts = append(workouts, &workout)
	}
	return workouts, nil
}

func (w *workout) GetWorkoutAssets(c *gin.Context, userID int64, planID int64) ([]*dto.WorkoutAsset, errcode.Error) {
	workoutDatas, err := w.workoutRepo.FindWorkoutAssets(userID, planID)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout repo", err)
	}
	workouts := make([]*dto.WorkoutAsset, 0)
	for _, workoutData := range workoutDatas {
		asset := dto.NewWorkoutAsset(workoutData)
		workouts = append(workouts, &asset)
	}
	return workouts, nil
}

func (w *workout) UpdateWorkout(c *gin.Context, workoutID int64, param *dto.UpdateWorkoutParam) (*dto.Workout, errcode.Error) {
	if err := w.workoutRepo.UpdateWorkoutByID(workoutID, &model.UpdateWorkoutParam{
		Name:      param.Name,
		Equipment: param.Equipment,
	}); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	var workout dto.Workout
	if err := w.workoutRepo.FindWorkoutByID(workoutID, &workout); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &workout, nil
}

func (w *workout) DeleteWorkout(c *gin.Context, workoutID int64) (*dto.WorkoutID, errcode.Error) {
	if err := w.workoutRepo.DeleteWorkoutByID(workoutID); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &dto.WorkoutID{ID: workoutID}, nil
}

func (w *workout) UploadWorkoutStartAudio(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error) {
	newAudioNamed, err := w.uploader.UploadWorkoutStartAudio(file, audioNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, w.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, w.errHandler.FileSizeError()
		}
		w.logger.Set(c, handler.Error, "Resource Handler", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	if err := w.workoutRepo.UpdateWorkoutByID(workoutID, &model.UpdateWorkoutParam{
		StartAudio: &newAudioNamed,
	}); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &dto.WorkoutAudio{Named: newAudioNamed}, nil
}

func (w *workout) UploadWorkoutEndAudio(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error) {
	newAudioNamed, err := w.uploader.UploadWorkoutEndAudio(file, audioNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, w.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, w.errHandler.FileSizeError()
		}
		w.logger.Set(c, handler.Error, "Resource Handler", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	if err := w.workoutRepo.UpdateWorkoutByID(workoutID, &model.UpdateWorkoutParam{
		EndAudio: &newAudioNamed,
	}); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &dto.WorkoutAudio{Named: newAudioNamed}, nil
}

func (w *workout) CreateWorkoutByTemplate(c *gin.Context, planID int64, name string, workoutTemplateID int64) (*dto.Workout, errcode.Error) {
	//驗證是否為相同的Course
	course1 := struct {
		ID int64 `gorm:"column:id"`
	}{}
	course2 := struct {
		ID int64 `gorm:"column:id"`
	}{}
	if err := w.courseRepo.FindCourseByPlanID(planID, &course1); err != nil {
		return nil, w.errHandler.Set(c, "course Repo", err)
	}
	if err := w.courseRepo.FindCourseByWorkoutID(workoutTemplateID, &course2); err != nil {
		return nil, w.errHandler.Set(c, "course Repo", err)
	}
	if course1.ID != course2.ID {
		return nil, w.errHandler.Set(c, "course Repo", errors.New(strconv.Itoa(errcode.PermissionDenied)))
	}
	//複製workout
	var template dto.Workout
	if err := w.workoutRepo.FindWorkoutByID(workoutTemplateID, &template); err != nil {
		return nil, w.errHandler.Set(c, "Workout Repo", err)
	}
	newWorkoutID, err := w.workoutRepo.CreateWorkout(planID, name)
	if err != nil {
		return nil, w.errHandler.Set(c, "Workout Repo", err)
	}
	if err := w.workoutRepo.UpdateWorkoutByID(newWorkoutID, &model.UpdateWorkoutParam{
		Equipment:  &template.Equipment,
		StartAudio: &template.StartAudio,
		EndAudio:   &template.EndAudio,
	}); err != nil {
		return nil, w.errHandler.Set(c, "Workout Repo", err)
	}
	//複製workout底下的sets
	entities, err := w.workoutSetRepo.FindWorkoutSetsByWorkoutID(workoutTemplateID, nil)
	if err != nil {
		return nil, w.errHandler.Set(c, "WorkoutSet Repo", err)
	}
	sets := make([]*entity.WorkoutSet, 0)
	for _, v := range entities {
		set := entity.WorkoutSet{
			WorkoutID:     newWorkoutID,
			Type:          v.Type,
			AutoNext:      v.AutoNext,
			StartAudio:    v.StartAudio,
			ProgressAudio: v.ProgressAudio,
			Remark:        v.Remark,
			Weight:        v.Weight,
			Reps:          v.Reps,
			Distance:      v.Distance,
			Duration:      v.Duration,
			Incline:       v.Incline,
			CreateAt:      time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt:      time.Now().Format("2006-01-02 15:04:05"),
		}
		if v.Action != nil {
			set.ActionID = &v.Action.ID
		}
		sets = append(sets, &set)
	}
	_, err = w.workoutSetRepo.CreateWorkoutSetsByWorkoutIDAndSets(newWorkoutID, sets)
	if err != nil {
		return nil, w.errHandler.Set(c, "WorkoutSet Repo", err)
	}
	//回傳此workout
	var workout dto.Workout
	if err := w.workoutRepo.FindWorkoutByID(newWorkoutID, &workout); err != nil {
		return nil, w.errHandler.Set(c, "Workout Repo", err)
	}
	return &workout, nil
}

func (w *workout) DeleteWorkoutStartAudio(c *gin.Context, workoutID int64) errcode.Error {
	var workout dto.Workout
	if err := w.workoutRepo.FindWorkoutByID(workoutID, &workout); err != nil {
		return w.errHandler.Set(c, "Workout Repo", err)
	}
	//移除檔案關聯
	var startAudio = ""
	if err := w.workoutRepo.UpdateWorkoutByID(workoutID, &model.UpdateWorkoutParam{
		StartAudio: &startAudio,
	}); err != nil {
		return w.errHandler.Set(c, "Workout Repo", err)
	}
	//查找是否有其他物件使用此物件
	count, err := w.workoutRepo.FindStartAudioCountByAudioName(workout.StartAudio)
	if err != nil {
		return w.errHandler.Set(c, "Workout Repo", err)
	}
	if count > 0 {
		return nil
	}
	//移除start_audio檔案
	if err := w.resHandler.DeleteWorkoutStartAudio(workout.StartAudio); err != nil {
		return w.errHandler.Set(c, "ResHandler", err)
	}
	return nil
}

func (w *workout) DeleteWorkoutEndAudio(c *gin.Context, workoutID int64) errcode.Error {
	var workout dto.Workout
	if err := w.workoutRepo.FindWorkoutByID(workoutID, &workout); err != nil {
		return w.errHandler.Set(c, "Workout Repo", err)
	}
	//移除檔案關聯
	var endAudio = ""
	if err := w.workoutRepo.UpdateWorkoutByID(workoutID, &model.UpdateWorkoutParam{
		EndAudio: &endAudio,
	}); err != nil {
		return w.errHandler.Set(c, "Workout Repo", err)
	}
	//查找是否有其他物件使用此物件
	count, err := w.workoutRepo.FindEndAudioCountByAudioName(workout.StartAudio)
	if err != nil {
		return w.errHandler.Set(c, "Workout Repo", err)
	}
	if count > 0 {
		return nil
	}
	//移除end_audio檔案
	if err := w.resHandler.DeleteWorkoutEndAudio(workout.EndAudio); err != nil {
		return w.errHandler.Set(c, "ResHandler", err)
	}
	return nil
}

func (w *workout) GetWorkoutStatus(c *gin.Context, workoutID int64) (global.CourseStatus, errcode.Error) {
	course := struct {
		CourseStatus int `json:"course_status"`
	}{}
	if err := w.courseRepo.FindCourseByWorkoutID(workoutID, &course); err != nil {
		return 0, w.errHandler.Set(c, "course repo", err)
	}
	return global.CourseStatus(course.CourseStatus), nil
}
