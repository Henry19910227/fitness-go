package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
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
	resHandler handler.Resource
	errHandler errcode.Handler
}

func NewAction(actionRepo repository.Action,
	courseRepo repository.Course,
	uploader handler.Uploader,
	resHandler handler.Resource,

	errHandler errcode.Handler) Action {
	return &action{actionRepo: actionRepo, courseRepo: courseRepo, uploader: uploader, resHandler: resHandler, errHandler: errHandler}
}

func (a *action) CreateAction(c *gin.Context, courseID int64, param *dto.CreateActionParam) (*dto.Action, errcode.Error) {
	 actionID, err := a.actionRepo.CreateAction(courseID, &model.CreateActionParam{
		 Name:      param.Name,
		 Type:      param.Type,
		 Category:  param.Category,
		 Body:      param.Body,
		 Equipment: param.Equipment,
		 Intro:     param.Intro,
	 })
	if err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	var action dto.Action
	if err := a.actionRepo.FindActionByID(actionID, &action); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return &action, nil
}

func (a *action) UpdateAction(c *gin.Context, actionID int64, param *dto.UpdateActionParam) (*dto.Action, errcode.Error) {
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Name: param.Name,
		Category: param.Category,
		Body: param.Body,
		Equipment: param.Equipment,
		Intro: param.Intro,
	}); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	var action dto.Action
	if err := a.actionRepo.FindActionByID(actionID, &action); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return &action, nil
}

func (a *action) SearchActions(c *gin.Context, courseID int64, param *dto.FindActionsParam) ([]*dto.Action, errcode.Error) {

	var sourceOpt []int
	if param.Source != nil {
		for _, item := range strings.Split(*param.Source, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				return nil, a.errHandler.Set(c, "strconv", err)
			}
			sourceOpt = append(sourceOpt, opt)
		}
	}

	var categoryOpt []int
	if param.Category != nil {
		for _, item := range strings.Split(*param.Category, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				return nil, a.errHandler.Set(c, "strconv", err)
			}
			categoryOpt = append(categoryOpt, opt)
		}
	}

	var bodyOpt []int
	if param.Body != nil {
		for _, item := range strings.Split(*param.Body, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				return nil, a.errHandler.Set(c, "strconv", err)
			}
			bodyOpt = append(bodyOpt, opt)
		}
	}

	var equipmentOpt []int
	if param.Equipment != nil {
		for _, item := range strings.Split(*param.Equipment, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				return nil, a.errHandler.Set(c, "strconv", err)
			}
			equipmentOpt = append(equipmentOpt, opt)
		}
	}

	var actions []*dto.Action
	if err := a.actionRepo.FindActionsByParam(courseID, &model.FindActionsParam{
		Name: param.Name,
		SourceOpt: &sourceOpt,
		CategoryOpt: &categoryOpt,
		BodyOpt: &bodyOpt,
		EquipmentOpt: &equipmentOpt,
	}, &actions); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return actions, nil
}

func (a *action) DeleteAction(c *gin.Context, actionID int64) (*dto.ActionID, errcode.Error) {
	if err := a.actionRepo.DeleteActionByID(actionID); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return &dto.ActionID{ID: actionID}, nil
}

func (a *action) UploadActionCover(c *gin.Context, actionID int64, coverNamed string, file multipart.File) (*dto.ActionCover, errcode.Error) {
	//上傳照片
	newImageNamed, err := a.uploader.UploadActionCover(file, coverNamed)
	if err != nil {
		return nil, a.errHandler.Set(c, "uploader", err)
	}
	//修改動作欄位
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Cover: &newImageNamed,
	}); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return &dto.ActionCover{Cover: newImageNamed}, nil
}

func (a *action) UploadActionVideo(c *gin.Context, actionID int64, videoNamed string, file multipart.File) (*dto.ActionVideo, errcode.Error) {
	//上傳影片
	newVideoNamed, err := a.uploader.UploadActionVideo(file, videoNamed)
	if err != nil {
		return nil, a.errHandler.Set(c, "uploader", err)
	}
	//修改動作欄位
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Video: &newVideoNamed,
	}); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return &dto.ActionVideo{Video: newVideoNamed}, nil
}

func (a *action) DeleteActionVideo(c *gin.Context, actionID int64) errcode.Error {
	var action dto.Action
	if err := a.actionRepo.FindActionByID(actionID, &action); err != nil {
		return a.errHandler.Set(c, "action repo", err)
	}
	var video = ""
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Video: &video,
	}); err != nil {
		return a.errHandler.Set(c, "action repo", err)
	}
	if err := a.resHandler.DeleteActionVideo(action.Video); err != nil {
		a.errHandler.Set(c, "resource handler", err)
	}
	return nil
}
