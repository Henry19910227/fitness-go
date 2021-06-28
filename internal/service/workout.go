package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strings"
)

type workout struct {
	workoutRepo repository.Workout
	planRepo repository.Plan
	courseRepo repository.Course
	uploader handler.Uploader
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewWorkout(workoutRepo repository.Workout, planRepo repository.Plan, courseRepo repository.Course, uploader handler.Uploader, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Workout {
	return &workout{workoutRepo: workoutRepo, planRepo: planRepo, courseRepo: courseRepo, uploader: uploader, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (w *workout) CreateWorkoutByToken(c *gin.Context, token string, planID int64, name string) (*workoutdto.WorkoutID, errcode.Error) {
	uid, err := w.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, w.errHandler.InvalidToken()
	}
	isExist, err := w.planRepo.CheckPlanExistByUID(uid, planID)
	if err != nil {
		return nil, w.errHandler.SystemError()
	}
	if !isExist {
		return nil, w.errHandler.PermissionDenied()
	}
	return w.CreateWorkout(c, planID, name)
}

func (w *workout) CreateWorkout(c *gin.Context, planID int64, name string) (*workoutdto.WorkoutID, errcode.Error) {
	workoutID, err := w.workoutRepo.CreateWorkout(planID, name)
	if err != nil {
		w.logger.Set(c, handler.Error, "CourseRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &workoutdto.WorkoutID{ID: workoutID}, nil
}

func (w *workout) GetWorkoutsByPlanID(c *gin.Context, planID int64) ([]*workoutdto.Workout, errcode.Error) {
	datas, err := w.workoutRepo.FindWorkoutsByPlanID(planID)
	if err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	workouts := make([]*workoutdto.Workout, 0)
	for _, data := range datas {
		workout := workoutdto.Workout{
			ID: data.ID,
			Name: data.Name,
			Equipment: data.Equipment,
			StartAudio: data.StartAudio,
			EndAudio: data.EndAudio,
			WorkoutSetCount: data.WorkoutSetCount,
		}
		workouts = append(workouts, &workout)
	}
	return workouts, nil
}

func (w *workout) UpdateWorkoutByToken(c *gin.Context, token string, workoutID int64, param *workoutdto.UpdateWorkoutParam) (*workoutdto.Workout, errcode.Error) {
	uid, err := w.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, w.errHandler.InvalidToken()
	}
	isExist, err := w.workoutRepo.CheckWorkoutExistByUID(uid, workoutID)
	if err != nil {
		return nil, w.errHandler.SystemError()
	}
	if !isExist {
		return nil, w.errHandler.PermissionDenied()
	}
	return w.UpdateWorkout(c, workoutID, param)
}

func (w *workout) UpdateWorkout(c *gin.Context, workoutID int64, param *workoutdto.UpdateWorkoutParam) (*workoutdto.Workout, errcode.Error) {
	if err := w.workoutRepo.UpdateWorkoutByID(workoutID, &model.UpdateWorkoutParam{
		Name: param.Name,
		Equipment: param.Equipment,
	}); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	var workout workoutdto.Workout
	if err := w.workoutRepo.FindWorkoutByID(workoutID, &workout); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &workout, nil
}

func (w *workout) DeleteWorkoutByToken(c *gin.Context, token string, workoutID int64) (*workoutdto.WorkoutID, errcode.Error) {
	uid, err := w.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, w.errHandler.InvalidToken()
	}
	isExist, err := w.workoutRepo.CheckWorkoutExistByUID(uid, workoutID)
	if err != nil {
		return nil, w.errHandler.SystemError()
	}
	if !isExist {
		return nil, w.errHandler.PermissionDenied()
	}
	return w.DeleteWorkout(c, workoutID)
}

func (w *workout) DeleteWorkout(c *gin.Context, workoutID int64) (*workoutdto.WorkoutID, errcode.Error) {
	status, err := w.courseRepo.FindCourseStatusByWorkoutID(workoutID)
	if err != nil {
		w.logger.Set(c, handler.Error, "CourseRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	if !(status == 1 || status == 4) {
		return nil, w.errHandler.PermissionDenied()
	}
	if err := w.workoutRepo.DeleteWorkoutByID(workoutID); err != nil {
		w.logger.Set(c, handler.Error, "WorkoutRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &workoutdto.WorkoutID{ID: workoutID}, nil
}

func (w *workout) UploadWorkoutStartAudioByToken(c *gin.Context, token string, workoutID int64, audioNamed string, file multipart.File) (*workoutdto.Audio, errcode.Error) {
	uid, err := w.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, w.errHandler.InvalidToken()
	}
	isExist, err := w.workoutRepo.CheckWorkoutExistByUID(uid, workoutID)
	if err != nil {
		return nil, w.errHandler.SystemError()
	}
	if !isExist {
		return nil, w.errHandler.PermissionDenied()
	}
	return w.UploadWorkoutStartAudioByID(c, workoutID, audioNamed, file)
}

func (w *workout) UploadWorkoutStartAudioByID(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*workoutdto.Audio, errcode.Error) {
	newAudioNamed, err := w.uploader.UploadWorkoutAudio(file, audioNamed)
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
	return &workoutdto.Audio{Named: newAudioNamed}, nil
}

func (w *workout) UploadWorkoutEndAudioByToken(c *gin.Context, token string, workoutID int64, audioNamed string, file multipart.File) (*workoutdto.Audio, errcode.Error) {
	uid, err := w.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, w.errHandler.InvalidToken()
	}
	isExist, err := w.workoutRepo.CheckWorkoutExistByUID(uid, workoutID)
	if err != nil {
		return nil, w.errHandler.SystemError()
	}
	if !isExist {
		return nil, w.errHandler.PermissionDenied()
	}
	return w.UploadWorkoutEndAudioByID(c, workoutID, audioNamed, file)
}

func (w *workout) UploadWorkoutEndAudioByID(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*workoutdto.Audio, errcode.Error) {
	newAudioNamed, err := w.uploader.UploadWorkoutAudio(file, audioNamed)
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
	return &workoutdto.Audio{Named: newAudioNamed}, nil
}