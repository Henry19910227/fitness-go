package bank_account

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetTrainerBankAccount(ctx *gin.Context)
	UpdateTrainerBankAccount(ctx *gin.Context)
}
