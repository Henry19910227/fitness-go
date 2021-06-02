package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Healthy struct {
}

func NewHealthy(router *gin.Engine) {
	healthy := &Healthy{}
	router.GET("/", healthy.Healthy)
}

func (h *Healthy) Healthy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "I feel good"})
}