package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type Favorite struct {
	Base
	favoriteService service.Favorite
	courseService   service.Course
}

func NewFavorite(baseGroup *gin.RouterGroup, favoriteService service.Favorite, courseService service.Course, userMidd midd.User, courseMidd midd.Course) {
	favorite := Favorite{favoriteService: favoriteService}
	baseGroup.POST("/favorite/course/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		favorite.CreateFavoriteCourse)
	baseGroup.POST("/favorite/trainer/:user_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		favorite.CreateFavoriteTrainer)
	baseGroup.POST("/favorite/action/:action_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		favorite.CreateFavoriteAction)
	baseGroup.DELETE("/favorite/course/:course_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		courseMidd.CourseStatusVerify(courseService.GetCourseStatus, []global.CourseStatus{global.Sale}),
		favorite.DeleteFavoriteCourse)
	baseGroup.DELETE("/favorite/trainer/:user_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		favorite.DeleteFavoriteTrainer)
	baseGroup.DELETE("/favorite/action/:action_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		favorite.DeleteFavoriteAction)
}

// CreateFavoriteCourse 新增收藏課表
// @Summary 新增收藏課表 (API已過時，更新為 /v2/favorite/course/{course_id} [POST])
// @Description 新增收藏課表
// @Tags Favorite_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult "新增成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/favorite/course/{course_id} [POST]
func (f *Favorite) CreateFavoriteCourse(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.favoriteService.CreateFavoriteCourse(c, uid, uri.CourseID); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}

// CreateFavoriteTrainer 收藏教練
// @Summary 收藏教練 (API已經過時，更新為 /v2/favorite/trainer/{user_id} [POST])
// @Description 收藏教練
// @Tags Favorite_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Success 200 {object} model.SuccessResult "新增成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/favorite/trainer/{user_id} [POST]
func (f *Favorite) CreateFavoriteTrainer(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.TrainerIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.favoriteService.CreateFavoriteTrainer(c, uid, uri.TrainerID); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}

// CreateFavoriteAction 收藏動作
// @Summary 收藏動作  (API已經過時，更新為 /v2/favorite/action/{action_id} [POST])
// @Description 收藏動作
// @Tags Favorite_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} model.SuccessResult "新增成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/favorite/action/{action_id} [POST]
func (f *Favorite) CreateFavoriteAction(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.favoriteService.CreateFavoriteAction(c, uid, uri.ActionID); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}

// DeleteFavoriteCourse 刪除收藏課表
// @Summary 刪除收藏課表 (API已經過時，更新為 /v2/favorite/course/{course_id} [DELETE])
// @Description 刪除收藏課表
// @Tags Favorite_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/favorite/course/{course_id} [DELETE]
func (f *Favorite) DeleteFavoriteCourse(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.CourseIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.favoriteService.DeleteFavoriteCourse(c, uid, uri.CourseID); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}

// DeleteFavoriteTrainer 刪除收藏教練
// @Summary 刪除收藏教練 (API已經過時，更新為 /v2/favorite/trainer/{user_id} [DELETE])
// @Description 刪除收藏教練
// @Tags Favorite_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/favorite/trainer/{user_id} [DELETE]
func (f *Favorite) DeleteFavoriteTrainer(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.TrainerIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.favoriteService.DeleteFavoriteTrainer(c, uid, uri.TrainerID); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}

// DeleteFavoriteAction 刪除收藏動作
// @Summary 刪除收藏動作 (API已經過時，更新為 /v2/favorite/action/{action_id} [DELETE])
// @Description 刪除收藏動作
// @Tags Favorite_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param action_id path int64 true "動作id"
// @Success 200 {object} model.SuccessResult "刪除成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /v1/favorite/action/{action_id} [DELETE]
func (f *Favorite) DeleteFavoriteAction(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.ActionIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.favoriteService.DeleteFavoriteAction(c, uid, uri.ActionID); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}
