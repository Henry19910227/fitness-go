package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fcm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/service/card"
	"github.com/Henry19910227/fitness-go/internal/v2/service/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_album"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_status_update_log"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	trainerService := trainer.NewService(db)
	trainerAlbumService := trainer_album.NewService(db)
	trainerStatusLogService := trainer_status_update_log.NewService(db)
	cardService := card.NewService(db)
	certService := certificate.NewService(db)
	bankAccountService := bank_account.NewService(db)
	avatarUploadTool := uploader.NewTrainerAvatarTool()
	albumUploadTool := uploader.NewTrainerAlbumTool()
	cardFrontUploadTool := uploader.NewCartFrontImageTool()
	cardBackUploadTool := uploader.NewCartBackImageTool()
	certUploadTool := uploader.NewCertificateImageTool()
	accountUploadTool := uploader.NewAccountImageTool()
	redisTool := redis.NewTool()
	fcmTool := fcm.NewTool()
	return New(trainerService, trainerAlbumService, trainerStatusLogService, cardService, certService, bankAccountService,
		avatarUploadTool, albumUploadTool, cardFrontUploadTool, cardBackUploadTool,
		certUploadTool, accountUploadTool, redisTool, fcmTool)
}
