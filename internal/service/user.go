package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mime/multipart"
	"strings"
)

type user struct {
	Base
	userRepo          repository.User
	trainerRepo       repository.Trainer
	subscribeInfoRepo repository.UserSubscribeInfo
	albumRepo         repository.TrainerAlbum
	certRepo          repository.Certificate
	uploader          handler.Uploader
	resHandler        handler.Resource
	logger            handler.Logger
	jwtTool           tool.JWT
	errHandler        errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer, subscribeInfoRepo repository.UserSubscribeInfo,
	albumRepo repository.TrainerAlbum, certRepo repository.Certificate,
	uploader handler.Uploader, resHandler handler.Resource, logger handler.Logger,
	jwtTool tool.JWT, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo,
		trainerRepo:       trainerRepo,
		subscribeInfoRepo: subscribeInfoRepo,
		albumRepo:         albumRepo,
		certRepo:          certRepo,
		uploader:          uploader,
		resHandler:        resHandler,
		logger:            logger,
		jwtTool:           jwtTool,
		errHandler:        errHandler}
}

func (u *user) UpdateUserByToken(c *gin.Context, token string, param *dto.UpdateUserParam) (*dto.User, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.UpdateUserByUID(c, uid, param)
}

func (u *user) UpdateUserByUID(c *gin.Context, uid int64, param *dto.UpdateUserParam) (*dto.User, errcode.Error) {
	//更新user
	if err := u.userRepo.UpdateUserByUID(nil, uid, &model.UpdateUserParam{
		Nickname:   param.Nickname,
		Sex:        param.Sex,
		Birthday:   param.Birthday,
		Height:     param.Height,
		Weight:     param.Weight,
		Experience: param.Experience,
		Target:     param.Target,
		UserStatus: param.UserStatus,
		Password:   param.Password,
	}); err != nil {
		//資料已存在
		if u.MysqlDuplicateEntry(err) {
			if strings.Contains(err.Error(), "nickname") {
				return nil, u.errHandler.Custom(9004, errors.New("重複的暱稱"))
			}
			return nil, u.errHandler.DataAlreadyExists()
		}
		//不明原因錯誤
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//查找user
	user, err := u.GetUserByUID(c, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) GetUserByUID(c *gin.Context, uid int64) (*dto.User, errcode.Error) {
	//查找user
	var user dto.User
	if err := u.userRepo.FindUserByUID(uid, &user); err != nil {
		//查無此資料
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, u.errHandler.DataNotFound()
		}
		//不明原因錯誤
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//獲取教練資訊
	data, err := u.trainerRepo.FindTrainer(uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Set(c, handler.Error, "TrainerRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	if data != nil {
		trainer := dto.NewTrainer(data)
		if err := u.albumRepo.FindAlbumPhotosByUID(user.ID, &trainer.TrainerAlbumPhotos); err != nil {
			return nil, u.errHandler.Set(c, "trainer album repo", err)
		}
		if err := u.certRepo.FindCertificatesByUID(user.ID, &trainer.Certificates); err != nil {
			return nil, u.errHandler.Set(c, "cer repo", err)
		}
		user.TrainerInfo = &trainer
	}
	//獲取訂閱資訊
	subscribeInfoData, err := u.subscribeInfoRepo.FindSubscribeInfo(uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, u.errHandler.Set(c, "user subscribe info repo", err)
	}
	if subscribeInfoData != nil {
		user.SubscribeInfo = &dto.UserSubscribeInfo{
			Status:      subscribeInfoData.Status,
			StartDate:   subscribeInfoData.StartDate,
			ExpiresDate: subscribeInfoData.ExpiresDate,
		}
	}
	return &user, nil
}

func (u *user) GetUserByToken(c *gin.Context, token string) (*dto.User, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.GetUserByUID(c, uid)
}

func (u *user) GetCMSUsers(c *gin.Context, param *dto.FinsCMSUsersParam, orderByParam *dto.OrderByParam, pagingParam *dto.PagingParam) ([]*dto.CMSUserSummary, *dto.Paging, errcode.Error) {
	//設置排序
	var orderBy *model.OrderBy
	if orderByParam != nil {
		orderBy = &model.OrderBy{
			OrderType: global.DESC,
			Field:     "create_at",
		}
		if orderByParam.OrderType != nil {
			orderBy.OrderType = global.OrderType(*orderByParam.OrderType)
		}
		if orderByParam.OrderField != nil {
			orderBy.Field = *orderByParam.OrderField
		}
	}
	//設置分頁
	var paging *model.PagingParam
	if pagingParam != nil {
		offset, limit := u.GetPagingIndex(pagingParam.Page, pagingParam.Size)
		paging = &model.PagingParam{
			Offset: offset,
			Limit:  limit,
		}
	}
	//獲取分頁資料
	users := make([]*dto.CMSUserSummary, 0)
	var totalCount int64
	if err := u.userRepo.FindUsers(&users, &totalCount, &model.FinsUsersParam{
		UserID:     param.UserID,
		Name:       param.Name,
		Email:      param.Email,
		UserStatus: param.UserStatus,
		UserType:   param.UserType,
	}, orderBy, paging); err != nil {
		return nil, nil, u.errHandler.Set(c, "user repo", err)
	}
	pagingResult := dto.Paging{
		TotalCount: int(totalCount),
		TotalPage:  u.GetTotalPage(int(totalCount), pagingParam.Size),
		Page:       pagingParam.Page,
		Size:       pagingParam.Size,
	}
	return users, &pagingResult, nil
}

func (u *user) GetCMSUser(c *gin.Context, userID int64) (*dto.CMSUser, errcode.Error) {
	var user dto.CMSUser
	if err := u.userRepo.FindUserByUID(userID, &user); err != nil {
		return nil, u.errHandler.Set(c, "user repo", err)
	}
	return &user, nil
}

func (u *user) UploadUserAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.UserAvatar, errcode.Error) {
	//上傳照片
	newImageNamed, err := u.uploader.UploadUserAvatar(imageFile, imageNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, u.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, u.errHandler.FileSizeError()
		}
		u.logger.Set(c, handler.Error, "Resource Handler", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//查詢用戶資訊
	var user struct {
		Avatar string `gorm:"column:avatar"`
	}
	if err := u.userRepo.FindUserByUID(uid, &user); err != nil {
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//修改教練資訊
	if err := u.userRepo.UpdateUserByUID(nil, uid, &model.UpdateUserParam{
		Avatar: &newImageNamed,
	}); err != nil {
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//刪除舊照片
	if len(user.Avatar) > 0 {
		if err := u.resHandler.DeleteUserAvatar(user.Avatar); err != nil {
			u.logger.Set(c, handler.Error, "ResHandler", u.errHandler.SystemError().Code(), err.Error())
		}
	}
	return &dto.UserAvatar{Avatar: newImageNamed}, nil
}

func (u *user) UploadUserAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*dto.UserAvatar, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.UploadUserAvatarByUID(c, uid, imageNamed, imageFile)
}
