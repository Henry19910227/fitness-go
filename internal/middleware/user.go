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
	jwtTool tool.JWT
	redisTool tool.Redis
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer, trainerAlbumRepo repository.TrainerAlbum, jwtTool tool.JWT, redisTool tool.Redis, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo, trainerRepo: trainerRepo, trainerAlbumRepo: trainerAlbumRepo, jwtTool: jwtTool, redisTool: redisTool, errHandler: errHandler}
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

func (u *user) TrainerAlbumPhotoLimit(count int) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			c.Abort()
			return
		}
		entities, err := u.trainerAlbumRepo.FindAlbumPhotoByUID(uid.(int64))
		if err != nil {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "trainer album repo", err))
			c.Abort()
			return
		}
		// 數量超過限制
		if len(entities) >= count {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "limit", errors.New(strconv.Itoa(errcode.FileCountError))))
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
