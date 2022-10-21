package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/trainer"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := trainer.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/trainer/avatar", http.Dir("./volumes/storage/trainer/avatar"))
	v2.GET("/trainer/profile", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerProfile)
	v2.GET("/trainer/:user_id", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainer)
	v2.GET("/favorite/trainers", midd.Verify([]global.Role{global.UserRole}), controller.GetFavoriteTrainers)
	v2.PATCH("/cms/trainer/:user_id/avatar", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSTrainerAvatar)
}
