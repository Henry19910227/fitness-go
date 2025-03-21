package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/bank_account"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := bank_account.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/trainer/account_image", http.Dir("./volumes/storage/trainer/account_image"))
	v2.GET("/trainer/bank_account", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerBankAccount)
	v2.PATCH("/trainer/bank_account", midd.Verify([]global.Role{global.UserRole}), controller.UpdateTrainerBankAccount)
}
