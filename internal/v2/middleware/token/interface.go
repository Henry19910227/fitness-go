package token

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Verify(roles []global.Role) gin.HandlerFunc
}