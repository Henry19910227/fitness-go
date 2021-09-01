package middleware

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Role int
const (
	UserRole Role = 1
	AdminRole = 2
)

type UserStatus int
const (
	UserActivity UserStatus = 1
	UserIllegal = 2
)

type TrainerStatus int
const (
	TrainerActivity TrainerStatus = 1
	TrainerReviewing = 2
	TrainerRevoke = 3
)

var (
	userTokenPrefix  = "fitness.user.token"
	adminTokenPrefix = "fitness.admin.token"
)


type user struct {
	Base
	userRepo repository.User
	trainerRepo repository.Trainer
	jwtTool tool.JWT
	redisTool tool.Redis
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer, jwtTool tool.JWT, redisTool tool.Redis, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo, trainerRepo: trainerRepo, jwtTool: jwtTool, redisTool: redisTool, errHandler: errHandler}
}

func (u *user) TokenPermission(roles []Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var header validator.TokenHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			u.JSONValidatorErrorResponse(c, err)
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
		if Role(role) == AdminRole {
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
		if !containRole(Role(role), roles) {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Set("role", role)
	}
}

func (u *user) UserStatusPermission(status []UserStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			return
		}
		if Role(role.(int)) == AdminRole {
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
		if !containUserStatus(status, UserStatus(user.UserStatus)) {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "user repo", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func (u *user) TrainerStatusPermission(status []TrainerStatus) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, isExists := c.Get("role")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			return
		}
		if Role(role.(int)) == AdminRole {
			return
		}
		uid, isExists := c.Get("uid")
		if !isExists {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
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
		}
		// 驗證是否包含所選的狀態
		if !containTrainerStatus(status, TrainerStatus(trainer.TrainerStatus)) {
			u.JSONErrorResponse(c, u.errHandler.Set(c, "jwt", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

func containUserStatus(items []UserStatus, target UserStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}

func containTrainerStatus(items []TrainerStatus, target TrainerStatus) bool {
	for _, v := range items {
		if target == v {
			return true
		}
	}
	return false
}

func containRole(role Role, roles []Role) bool {
	for _, v := range roles {
		if role == v {
			return true
		}
	}
	return false
}
