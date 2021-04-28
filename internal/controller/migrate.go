package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Migrate struct {
	Base
	migService service.Migrate
}

func NewMigrate(baseGroup *gin.RouterGroup, migService service.Migrate) {

	migrate := &Migrate{migService: migService}
	migrateGroup := baseGroup.Group("/migrate")
	migrateGroup.PUT("/up", migrate.UpToLatest)
	migrateGroup.PUT("/up/:step", migrate.UpStep)
	migrateGroup.PUT("/down", migrate.DownToOldest)
	migrateGroup.PUT("/down/:step", migrate.DownStep)
	migrateGroup.PUT("/force/:version", migrate.Force)
	migrateGroup.PUT("/version/:version", migrate.Migrate)
	migrateGroup.GET("/version", migrate.Version)
}

// Version 獲取當前 Schema 版本
// @Summary 獲取當前 Schema 版本
// @Description 獲取當前 Schema 版本
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/version [GET]
func (m *Migrate) Version(c *gin.Context) {
	version, isDirty, err := m.migService.Version()
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, gin.H{"version": version, "dirty": isDirty}, "success!")
}

// UpToLatest 將 Schema 升至最新版本
// @Summary 將 Schema 升至最新版本
// @Description 將 Schema 升至最新版本
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/up [PUT]
func (m *Migrate) UpToLatest(c *gin.Context) {
	version, isDirty, err := m.migService.Up()
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, gin.H{"version": version, "dirty": isDirty}, "success up migrate!")
}

// UpStep 將 Schema 升級N個版本
// @Summary 將 Schema 升級N個版本
// @Description 將 Schema 升級N個版本
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Param step path int true "版本跨度"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/up/{step} [PUT]
func (m *Migrate) UpStep(c *gin.Context) {
	var uri validator.MigrateStepUri
	if err := c.ShouldBindUri(&uri); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	version, isDirty, err := m.migService.UpStep(uri.Step)
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, gin.H{"version": version, "dirty": isDirty}, "success up migrate!")
}

// DownToOldest 回滾 Schema 至初始版本
// @Summary 回滾 Schema 至初始版本
// @Description 回滾 Schema 至初始版本
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/down [PUT]
func (m *Migrate) DownToOldest(c *gin.Context) {
	err := m.migService.Down()
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, nil, "success down migrate!")
}

// DownStep 將 Schema 回滾N個版本
// @Summary 將 Schema 回滾N個版本
// @Description 將 Schema 回滾N個版本
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Param step path int true "版本跨度"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/down/{step} [PUT]
func (m *Migrate) DownStep(c *gin.Context) {
	var uri validator.MigrateStepUri
	if err := c.ShouldBindUri(&uri); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	err := m.migService.DownStep(uri.Step)
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	version, isDirty, err := m.migService.Version()
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, gin.H{"version": version, "dirty": isDirty}, "success down migrate!")
}

// Force 修正 Schema 版本並解除錯誤狀態
// @Summary 修正 Schema 版本並解除錯誤狀態
// @Description Schema 升級時遇到錯誤時的操作模式
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Param version path int true "版本號"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/force/{version} [PUT]
func (m *Migrate) Force(c *gin.Context) {
	var uri validator.MigrateVersionUri
	if err := c.ShouldBindUri(&uri); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	version, isDirty, err := m.migService.Force(uri.Version)
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, gin.H{"version": version, "dirty": isDirty}, "success force migrate!")
}

// Migrate 升級至指定 Schema 版本
// @Summary 升級至指定 Schema 版本
// @Description 升級至指定 Schema 版本
// @Tags Migrate
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Param version path int true "版本號"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /migrate/version/{version} [PUT]
func (m *Migrate) Migrate(c *gin.Context) {
	var uri validator.MigrateVersionUri
	if err := c.ShouldBindUri(&uri); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	version, isDirty, err := m.migService.Migrate(uint(uri.Version))
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, gin.H{"version": version, "dirty": isDirty}, "success migrate!")
}
