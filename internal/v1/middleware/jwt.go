package middleware

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuth(jwtTool tool.JWT, e errcode.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		//驗證token不得為空
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "token not null"})
			c.Abort()
			return
		}
		//uid, err := jwtTool.GetIDByToken(token)
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"code": e.InvalidToken().Code(), "data": nil, "msg": e.InvalidToken().Msg()})
		//	c.Abort()
		//	return
		//}
		//role, err := jwtTool.GetRoleByToken(token)
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"code": e.InvalidToken().Code(), "data": nil, "msg": e.InvalidToken().Msg()})
		//	c.Abort()
		//	return
		//}

	}
}

func UserJWT(ssoHandler handler.SSO, e errcode.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		//驗證token不得為空
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "token not null"})
			c.Abort()
			return
		}
		//驗證token有效性
		if err := ssoHandler.VerifyUserToken(token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": e.InvalidToken().Code(), "data": nil, "msg": e.InvalidToken().Msg()})
			c.Abort()
			return
		}
	}
}

func AdminLV1JWT(ssoHandler handler.SSO, e errcode.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "token not null"})
			c.Abort()
			return
		}
		if err := ssoHandler.VerifyLV1AdminToken(token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": e.InvalidToken().Code(), "data": nil, "msg": e.InvalidToken().Msg()})
			c.Abort()
			return
		}
	}
}

func AdminLV2JWT(ssoHandler handler.SSO, e errcode.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "token not null"})
			c.Abort()
			return
		}
		if err := ssoHandler.VerifyLV2AdminToken(token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": e.InvalidToken().Code(), "data": nil, "msg": e.InvalidToken().Msg()})
			c.Abort()
			return
		}
	}
}
