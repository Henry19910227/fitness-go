package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/model/admindata"
	"github.com/gin-gonic/gin"
)

type Migrate interface {
	Version() (uint, bool, errcode.Error)
	Up() (uint, bool, errcode.Error)
	UpStep(step int) (uint, bool, errcode.Error)
	Down() errcode.Error
	DownStep(step int) errcode.Error
	Force(version int) (uint, bool, errcode.Error)
	Migrate(version uint) (uint, bool, errcode.Error)
}

type Swagger interface {
	WrapHandler() gin.HandlerFunc
}

type Login interface {
	Logout(c *gin.Context, token string) errcode.Error
	LoginForAdmin(c *gin.Context, email string, password string) (*admindata.Admin, string, errcode.Error)
	LogoutForAdmin(c *gin.Context, token string) errcode.Error
}
