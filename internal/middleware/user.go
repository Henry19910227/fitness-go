package middleware

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	userTokenPrefix  = "fitness.user.token"
	adminTokenPrefix = "fitness.admin.token"
)


type user struct {
	Base
	userRepo repository.User
	trainerRepo repository.Trainer
	trainerAlbumRepo repository.TrainerAlbum
	cerRepo repository.Certificate
	jwtTool tool.JWT
	redisTool tool.Redis
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer, trainerAlbumRepo repository.TrainerAlbum, cerRepo repository.Certificate, jwtTool tool.JWT, redisTool tool.Redis, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo, trainerRepo: trainerRepo, trainerAlbumRepo: trainerAlbumRepo, cerRepo: cerRepo, jwtTool: jwtTool, redisTool: redisTool, errHandler: errHandler}
}

func (u *user) TokenPermission(roles []global.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var header validator.TokenHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			u.JSONValidatorErrorResponse(c, err)
			c.Abort()
			return
		}
		uid, err := u.jwtTool.GetIDByToken(header.Token)
		if err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		role, err := u.jwtTool.GetRoleByToken(header.Token)
		if err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		// 驗證當前緩存的token是否過期
		key := userTokenPrefix + "." + strconv.Itoa(int(uid))
		if global.Role(role) == global.AdminRole {
			key = adminTokenPrefix + "." + strconv.Itoa(int(uid))
		}
		currentToken, err := u.redisTool.Get(key)
		if err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		if header.Token != currentToken {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		// 驗證是否包含所選的身份
		if !containRole(global.Role(role), roles) {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Set("role", role)
	}
}

func (u *user) UserStatusPermission(status []global.UserStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			return
		}
		user := struct {
			UserStatus int `gorm:"column:user_status"`
		}{}
		if err := u.userRepo.FindUserByUID(uid.(int64), &user); err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "user repo", err))
			c.Abort()
			return
		}
		// 驗證是否包含所選的狀態
		if !containUserStatus(status, global.UserStatus(user.UserStatus)) {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "user repo", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (u *user) TrainerStatusPermission(status []global.TrainerStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			c.Abort()
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			c.Abort()
			return
		}
		trainer := struct {
			UserID int64 `gorm:"column:user_id"`
			TrainerStatus int `gorm:"column:trainer_status"`
		}{}
		if err := u.trainerRepo.FindTrainerByUID(uid.(int64), &trainer); err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", err))
			c.Abort()
			return
		}
		// 此人不是教練
		if uid == 0 {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
		// 驗證是否包含所選的狀態
		if !containTrainerStatus(status, global.TrainerStatus(trainer.TrainerStatus)) {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (u *user) TrainerAlbumPhotoLimit(currentCount func(c *gin.Context, uid int64) (int, errcode.Error), createCount, deleteCount func(c *gin.Context) int, limitCount int) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			return
		}
		var currentCountValue int
		if currentCount != nil {
			count, err := currentCount(c, uid.(int64))
			if err != nil {
				c.Abort()
				u.JSONErrorResponse(c, err)
				return
			}
			currentCountValue = count
		}
		var createCountValue int
		if createCount != nil {
			createCountValue = createCount(c)
		}
		var deleteCountValue int
		if deleteCount != nil {
			deleteCountValue = deleteCount(c)
		}
		if (currentCountValue + createCountValue - deleteCountValue) > limitCount {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "limit", errors.New(strconv.Itoa(errcode.FileCountError))))
			c.Abort()
			return
		}
	}
}

func (u *user) CertificateCreatorVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		var uri validator.CertificateIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			u.JSONValidatorErrorResponse(c, err)
			c.Abort()
			return
		}
		certificate := struct {
			UserID int64 `gorm:"column:user_id"`
		}{}
		if err := u.cerRepo.FindCertificate(uri.CerID, &certificate); err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		if certificate.UserID != uid {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "verify", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (u *user) TrainerAlbumPhotoCreatorVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		if global.Role(role.(int)) == global.AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.InvalidToken))))
			c.Abort()
			return
		}
		var uri validator.TrainerAlbumPhotoIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			u.JSONValidatorErrorResponse(c, err)
			c.Abort()
			return
		}
		photo := struct {
			UserID int64 `gorm:"column:user_id"`
		}{}
		if err := u.trainerAlbumRepo.FindAlbumPhotoByID(uri.PhotoID, &photo); err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", err))
			c.Abort()
			return
		}
		if photo.UserID != uid {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "verify", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func containUserStatus(items []global.UserStatus, target global.UserStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}

func containTrainerStatus(items []global.TrainerStatus, target global.TrainerStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}

func containRole(role global.Role, roles []global.Role) bool {
	for _, v := range roles {
		if role == v {
			return true
		}
	}
	return false
}
