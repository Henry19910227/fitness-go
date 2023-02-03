package migrate

import "github.com/gin-gonic/gin"

type Controller interface {
	MigrateUpToLatest(ctx *gin.Context)
}
