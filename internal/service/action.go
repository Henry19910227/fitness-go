package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type action struct {
	Base
	actionRepo repository.Action
	courseRepo repository.Course
	uploader   handler.Uploader
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
	//生成Cover名稱
	var cover *string
	if param.Cover != nil {
		coverImageNamed, err := a.uploader.GenerateNewImageName(param.Cover.FileNamed)
		if err != nil {
			return nil, a.errHandler.Set(c, "uploader", err)
		}
		cover = &coverImageNamed
		param.Cover.FileNamed = coverImageNamed
	}
	//生成Video名稱
	var video *string
	if param.Video != nil {
		videoNamed, err := a.uploader.GenerateNewVideoName(param.Video.FileNamed)
		if err != nil {
			return nil, a.errHandler.Set(c, "uploader", err)
		}
		video = &videoNamed
		param.Video.FileNamed = videoNamed
	}
	//創建動作
	actionID, err := a.actionRepo.CreateAction(courseID, &model.CreateActionParam{
		Name:      param.Name,
		Type:      param.Type,
		Category:  param.Category,
		Body:      param.Body,
		Equipment: param.Equipment,
		Intro:     param.Intro,
		Video:     video,
		Cover:     cover,
	})
	if err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	//儲存動作封面照
	if param.Cover != nil {
		err = a.uploader.UploadActionCover(param.Cover.Data, param.Cover.FileNamed)
		if err != nil {
			a.errHandler.Set(c, "uploader", err)
		}
	}
	//儲存動作影片
	if param.Video != nil {
		err = a.uploader.UploadActionVideo(param.Video.Data, param.Video.FileNamed)
		if err != nil {
			a.errHandler.Set(c, "uploader", err)
		}
	}
	data, err := a.actionRepo.FindActionByID(actionID)
	if err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	action := dto.NewAction(data)
	return &action, nil
}

func (a *action) UpdateAction(c *gin.Context, actionID int64, param *dto.UpdateActionParam) (*dto.Action, errcode.Error) {
	//生成Cover名稱
	var cover *string
	if param.Cover != nil {
		coverImageNamed, err := a.uploader.GenerateNewImageName(param.Cover.FileNamed)
		if err != nil {
			return nil, a.errHandler.Set(c, "uploader", err)
		}
		cover = &coverImageNamed
		param.Cover.FileNamed = coverImageNamed
	}
	//生成Video名稱
	var video *string
	if param.Video != nil {
		videoNamed, err := a.uploader.GenerateNewVideoName(param.Video.FileNamed)
		if err != nil {
			return nil, a.errHandler.Set(c, "uploader", err)
		}
		video = &videoNamed
		param.Video.FileNamed = videoNamed
	}
	//查詢更新前的動作資料
	data, err := a.actionRepo.FindActionByID(actionID)
	if err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	oldAction := dto.NewAction(data)
	//更新動作
	if err := a.actionRepo.UpdateActionByID(actionID, &model.UpdateActionParam{
		Name:      param.Name,
		Category:  param.Category,
		Body:      param.Body,
		Equipment: param.Equipment,
		Intro:     param.Intro,
		Cover:     cover,
		Video:     video,
	}); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	//處理動作封面照
	if param.Cover != nil {
		//刪除舊的動作封面照
		if len(oldAction.Cover) > 0 {
			if err := a.resHandler.DeleteActionCover(oldAction.Cover); err != nil {
				a.errHandler.Set(c, "res handler", err)
			}
		}
		//修改新的動作封面照
		if err := a.uploader.UploadActionCover(param.Cover.Data, param.Cover.FileNamed); err != nil {
			a.errHandler.Set(c, "uploader", err)
		}
	}
	//處理動作影片
	if param.Video != nil {
		//刪除舊的動作影片
		if len(oldAction.Video) > 0 {
			if err := a.resHandler.DeleteActionVideo(oldAction.Video); err != nil {
				a.errHandler.Set(c, "res handler", err)
			}
		}
		//上傳新的動作影片
		if err := a.uploader.UploadActionVideo(param.Video.Data, param.Video.FileNamed); err != nil {
			a.errHandler.Set(c, "uploader", err)
		}
	}
	data, err = a.actionRepo.FindActionByID(actionID)
	if err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	action := dto.NewAction(data)
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
	datas, err := a.actionRepo.FindActionsByParam(courseID, &model.FindActionsParam{
		Name:         param.Name,
		SourceOpt:    &sourceOpt,
		CategoryOpt:  &categoryOpt,
		BodyOpt:      &bodyOpt,
		EquipmentOpt: &equipmentOpt,
	})
	if err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	var actions []*dto.Action
	for _, data := range datas {
		action := dto.NewAction(data)
		actions = append(actions, &action)
	}
	return actions, nil
}

func (a *action) DeleteAction(c *gin.Context, actionID int64) (*dto.ActionID, errcode.Error) {
	if err := a.actionRepo.DeleteActionByID(actionID); err != nil {
		return nil, a.errHandler.Set(c, "action repo", err)
	}
	return &dto.ActionID{ID: actionID}, nil
}

func (a *action) DeleteActionVideo(c *gin.Context, actionID int64) errcode.Error {
	data, err := a.actionRepo.FindActionByID(actionID)
	if err != nil {
		return a.errHandler.Set(c, "action repo", err)
	}
	action := dto.NewAction(data)
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
