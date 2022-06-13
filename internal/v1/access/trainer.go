package access

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
)

type trainer struct {
	trainerRepo repository.Trainer
	logger     handler.Logger
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewTrainer(trainerRepo repository.Trainer, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Trainer {
	return &trainer{trainerRepo: trainerRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (t *trainer) StatusVerify(c *gin.Context, token string) errcode.Error {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return t.errHandler.InvalidToken()
	}
	var trainer struct{
		UserID int64 `gorm:"column:user_id"`
		TrainerStatus int `gorm:"column:trainer_status"`
	}
	if err := t.trainerRepo.FindTrainerEntity(uid, &trainer); err != nil{
		t.logger.Set(c, handler.Error, "TrainerRepo", t.errHandler.SystemError().Code(), err.Error())
		return t.errHandler.SystemError()
	}
	if trainer.UserID == 0 {
		return t.errHandler.PermissionDenied()
	}
	if trainer.TrainerStatus == 3 {
		return t.errHandler.PermissionDenied()
	}
	return nil
}

