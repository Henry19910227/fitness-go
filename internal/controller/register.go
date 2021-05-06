package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Base
	regService service.Register
}

func NewRegister(baseGroup *gin.RouterGroup, regService service.Register)  {
	register := &Register{regService: regService}
	baseGroup.POST("/register/email", register.RegisterForEmail)
}

func (r *Register) RegisterForEmail(c *gin.Context)  {

}
