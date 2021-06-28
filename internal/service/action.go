package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/actiondto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"strings"
)

type action struct {
	Base
	actionRepo repository.Action
	courseRepo repository.Course
	uploader  handler.Uploader
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewAction(actionRepo repository.Action,
	courseRepo repository.Course,
	uploader handler.Uploader,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Action {
	return &action{actionRepo: actionRepo, courseRepo: courseRepo, uploader: uploader, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (a *action) CreateActionByToken(c *gin.Context, token string, courseID int64, param *actiondto.CreateActionParam) (*actiondto.Action, errcode.Error) {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, a.errHandler.InvalidToken()
	}
	isExist, err := a.courseRepo.CheckCourseExistByIDAndUID(courseID, uid)
	if err != nil {
		return nil, a.errHandler.SystemError()
	}
	if !isExist {
		return nil, a.errHandler.PermissionDenied()
	}
	return a.CreateAction(c, courseID, param)
}

func (a *action) CreateAction(c *gin.Context, courseID int64, param *actiondto.CreateActionParam) (*actiondto.Action, errcode.Error) {
	 actionID, err := a.actionRepo.CreateAction(courseID, &model.CreateActionParam{
		 Name:      param.Name,
		 Type:      param.Type,
		 Category:  param.Category,
		 Body:      param.Body,
		 Equipment: param.Equipment,
		 Intro:     param.Intro,
	 })
	if err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	var action actiondto.Action
	if err := a.actionRepo.FindActionByID(actionID, &action); err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	return &action, nil
}

func (a *action) UpdateActionByToken(c *gin.Context, token string, actionID int64, param *actiondto.UpdateActionParam) (*actiondto.Action, errcode.Error) {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, a.errHandler.InvalidToken()
	}
	isExist, err := a.actionRepo.CheckActionExistByUID(uid, actionID)
	if err != nil {
		return nil, a.errHandler.SystemError()
	}
	if !isExist {
		return nil, a.errHandler.PermissionDenied()
	}
	return a.UpdateAction(c, actionID, param)
}

func (a *action) UpdateAction(c *gin.Context, actionID int64, param *actiondto.UpdateActionParam) (*actiondto.Action, errcode.Error) {
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Name: param.Name,
		Category: param.Category,
		Body: param.Body,
		Equipment: param.Equipment,
		Intro: param.Intro,
	}); err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	var action actiondto.Action
	if err := a.actionRepo.FindActionByID(actionID, &action); err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	return &action, nil
}

func (a *action) SearchActionsByToken(c *gin.Context, token string, courseID int64, param *actiondto.FindActionsParam) ([]*actiondto.Action, errcode.Error) {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, a.errHandler.InvalidToken()
	}
	isExist, err := a.courseRepo.CheckCourseExistByIDAndUID(courseID, uid)
	if err != nil {
		return nil, a.errHandler.SystemError()
	}
	if !isExist {
		return nil, a.errHandler.PermissionDenied()
	}
	return a.SearchActions(c, courseID, param)
}

func (a *action) SearchActions(c *gin.Context, courseID int64, param *actiondto.FindActionsParam) ([]*actiondto.Action, errcode.Error) {

	var sourceOpt []int
	if param.Source != nil {
		for _, item := range strings.Split(*param.Source, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				a.logger.Set(c, handler.Error, "strconv", a.errHandler.SystemError().Code(), err.Error())
				return nil, a.errHandler.SystemError()
			}
			sourceOpt = append(sourceOpt, opt)
		}
	}

	var categoryOpt []int
	if param.Category != nil {
		for _, item := range strings.Split(*param.Category, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				a.logger.Set(c, handler.Error, "strconv", a.errHandler.SystemError().Code(), err.Error())
				return nil, a.errHandler.SystemError()
			}
			categoryOpt = append(categoryOpt, opt)
		}
	}

	var bodyOpt []int
	if param.Body != nil {
		for _, item := range strings.Split(*param.Body, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				a.logger.Set(c, handler.Error, "strconv", a.errHandler.SystemError().Code(), err.Error())
				return nil, a.errHandler.SystemError()
			}
			bodyOpt = append(bodyOpt, opt)
		}
	}

	var equipmentOpt []int
	if param.Equipment != nil {
		for _, item := range strings.Split(*param.Equipment, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				a.logger.Set(c, handler.Error, "strconv", a.errHandler.SystemError().Code(), err.Error())
				return nil, a.errHandler.SystemError()
			}
			equipmentOpt = append(equipmentOpt, opt)
		}
	}

	var actions []*actiondto.Action
	if err := a.actionRepo.FindActionsByParam(courseID, &model.FindActionsParam{
		Name: param.Name,
		SourceOpt: &sourceOpt,
		CategoryOpt: &categoryOpt,
		BodyOpt: &bodyOpt,
		EquipmentOpt: &equipmentOpt,
	}, &actions); err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	return actions, nil
}

func (a *action) DeleteActionByToken(c *gin.Context, token string, actionID int64) (*actiondto.ActionID, errcode.Error) {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, a.errHandler.InvalidToken()
	}
	isExist, err := a.actionRepo.CheckActionExistByUID(uid, actionID)
	if err != nil {
		return nil, a.errHandler.SystemError()
	}
	if !isExist {
		return nil, a.errHandler.PermissionDenied()
	}
	return a.DeleteAction(c, actionID)
}

func (a *action) DeleteAction(c *gin.Context, actionID int64) (*actiondto.ActionID, errcode.Error) {
	if err := a.actionRepo.DeleteActionByID(actionID); err != nil {
		if strings.Contains(err.Error(), "9006") {
			a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.PermissionDenied().Code(), err.Error())
			return nil, a.errHandler.PermissionDenied()
		}
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	return &actiondto.ActionID{ID: actionID}, nil
}

func (a *action) UploadActionCoverByToken(c *gin.Context, token string, actionID int64, coverNamed string, file multipart.File) (*actiondto.ActionCover, errcode.Error) {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, a.errHandler.InvalidToken()
	}
	isExist, err := a.actionRepo.CheckActionExistByUID(uid, actionID)
	if err != nil {
		return nil, a.errHandler.SystemError()
	}
	if !isExist {
		return nil, a.errHandler.PermissionDenied()
	}
	return a.UploadActionCover(c, actionID, coverNamed, file)
}

func (a *action) UploadActionCover(c *gin.Context, actionID int64, coverNamed string, file multipart.File) (*actiondto.ActionCover, errcode.Error) {
	//上傳照片
	newImageNamed, err := a.uploader.UploadActionCover(file, coverNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, a.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, a.errHandler.FileSizeError()
		}
		a.logger.Set(c, handler.Error, "Resource Handler", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	//修改動作欄位
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Cover: &newImageNamed,
	}); err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	return &actiondto.ActionCover{Cover: newImageNamed}, nil
}

func (a *action) UploadActionVideoByToken(c *gin.Context, token string, actionID int64, videoNamed string, file multipart.File) (*actiondto.ActionVideo, errcode.Error) {
	uid, err := a.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, a.errHandler.InvalidToken()
	}
	isExist, err := a.actionRepo.CheckActionExistByUID(uid, actionID)
	if err != nil {
		return nil, a.errHandler.SystemError()
	}
	if !isExist {
		return nil, a.errHandler.PermissionDenied()
	}
	return a.UploadActionVideo(c, actionID, videoNamed, file)
}

func (a *action) UploadActionVideo(c *gin.Context, actionID int64, videoNamed string, file multipart.File) (*actiondto.ActionVideo, errcode.Error) {
	//上傳影片
	newVideoNamed, err := a.uploader.UploadActionVideo(file, videoNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, a.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, a.errHandler.FileSizeError()
		}
		a.logger.Set(c, handler.Error, "Resource Handler", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	//修改動作欄位
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Video: &newVideoNamed,
	}); err != nil {
		a.logger.Set(c, handler.Error, "ActionRepo", a.errHandler.SystemError().Code(), err.Error())
		return nil, a.errHandler.SystemError()
	}
	return &actiondto.ActionVideo{Video: newVideoNamed}, nil
}
