package token

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	output "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	userTokenPrefix  = "fitness.user.token"
	adminTokenPrefix = "fitness.admin.token"
)

type middleware struct {
	jwtTool   tool.JWT
	redisTool tool.Redis
}

func NewMiddleware(jwtTool tool.JWT, redisTool tool.Redis) Middleware {
	return &middleware{jwtTool: jwtTool, redisTool: redisTool}
}

func (m *middleware) Verify(roles []global.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header validator.TokenHeader
		if err := ctx.ShouldBindHeader(&header); err != nil {
			ctx.JSON(http.StatusBadRequest, output.BadRequest(util.PointerString(err.Error())))
			ctx.Abort()
			return
		}
		uid, err := m.jwtTool.GetIDByToken(header.Token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, output.InvalidToken())
			ctx.Abort()
			return
		}
		role, err := m.jwtTool.GetRoleByToken(header.Token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, output.InvalidToken())
			ctx.Abort()
			return
		}
		// 驗證當前緩存的token是否過期
		key := userTokenPrefix + "." + strconv.Itoa(int(uid))
		if global.Role(role) == global.AdminRole {
			key = adminTokenPrefix + "." + strconv.Itoa(int(uid))
		}
		currentToken, err := m.redisTool.Get(key)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, output.InvalidToken())
			ctx.Abort()
			return
		}
		if header.Token != currentToken {
			ctx.JSON(http.StatusBadRequest, output.InvalidToken())
			ctx.Abort()
			return
		}
		// 驗證是否包含所選的身份
		if !containRole(global.Role(role), roles) {
			ctx.JSON(http.StatusBadRequest, output.PermissionDenied())
			ctx.Abort()
			return
		}
		ctx.Set("uid", uid)
		ctx.Set("role", role)
		ctx.Next()
	}
}

func containRole(role global.Role, roles []global.Role) bool {
	for _, v := range roles {
		if role == v {
			return true
		}
	}
	return false
}
