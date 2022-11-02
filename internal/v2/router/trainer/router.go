package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := trainer.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/trainer/avatar", http.Dir("./volumes/storage/trainer/avatar"))
	v2.StaticFS("/resource/trainer/card_front_image", http.Dir("./volumes/storage/trainer/card_front_image"))
	v2.StaticFS("/resource/trainer/card_back_image", http.Dir("./volumes/storage/trainer/card_back_image"))
	v2.StaticFS("/resource/trainer/album", http.Dir("./volumes/storage/trainer/album"))
	v2.StaticFS("/resource/trainer/certificate", http.Dir("./volumes/storage/trainer/certificate"))

	v2.POST("/trainer", middleware.Transaction(orm.Shared().DB()) , midd.Verify([]global.Role{global.UserRole}), controller.CreateTrainer)
	v2.GET("/trainer/profile", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerProfile)
	v2.GET("/store/trainer/:user_id", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreTrainer)
	v2.GET("/store/trainers", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreTrainers)
	v2.GET("/favorite/trainers", midd.Verify([]global.Role{global.UserRole}), controller.GetFavoriteTrainers)
	v2.PATCH("/cms/trainer/:user_id/avatar", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSTrainerAvatar)
}
