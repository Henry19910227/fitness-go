package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type CMSUser struct {
	Base
	userService service.User
}

func NewCMSUser(baseGroup *gin.RouterGroup, userService service.User, userMiddleware middleware.User) {
	cms := &CMSUser{userService: userService}
	baseGroup.GET("/cms/users",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetUsers)

	baseGroup.GET("/cms/user/:user_id",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.GetUser)

	baseGroup.PATCH("/cms/user/:user_id",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.UpdateUser)
}

// GetUsers 獲取用戶列表
// @Summary 獲取用戶列表
// @Description 獲取用戶列表
// @Tags CMS/User
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id query int64 false "用戶ID"
// @Param nickname query string false "用戶名稱(1~40字元)"
// @Param email query string false "用戶Email"
// @Param user_status query string false "用戶狀態 (1:正常/2:違規/3:刪除)"
// @Param user_type query string false "用戶類型 (1:一般用戶/2:訂閱用戶)"
// @Param order_field query string false "排序欄位 (create_at:創建時間)"
// @Param order_type query string false "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=[]dto.CMSUserSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /cms/users [GET]
func (u *CMSUser) GetUsers(c *gin.Context) {
	var form validator.CMSGetUsersQuery
	if err := c.ShouldBind(&form); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var orderByQuery validator.OrderByQuery
	if err := c.ShouldBind(&orderByQuery); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var pagingQuery validator.PagingQuery
	if err := c.ShouldBind(&pagingQuery); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	users, paging, err := u.userService.GetCMSUsers(c, &dto.FinsCMSUsersParam{
		UserID:     form.UserID,
		Name:       form.Nickname,
		Email:      form.Email,
		UserStatus: form.UserStatus,
		UserType:   form.UserType,
	}, &dto.OrderByParam{
		OrderField: orderByQuery.OrderField,
		OrderType:  orderByQuery.OrderType,
	}, &dto.PagingParam{
		Page: pagingQuery.Page,
		Size: pagingQuery.Size,
	})
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessPagingResponse(c, users, paging, "success!")
}

// GetUser 取得用戶詳細資訊
// @Summary 取得用戶詳細資訊
// @Description 取得用戶詳細資訊
// @Tags CMS/User
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "用戶id"
// @Success 200 {object} model.SuccessResult{data=dto.CMSUser} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /cms/user/{user_id} [GET]
func (u *CMSUser) GetUser(c *gin.Context) {
	var uri validator.UserIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	user, err := u.userService.GetCMSUser(c, *uri.UserID)
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessResponse(c, user, "success!")
}

// UpdateUser 更新用戶資訊
// @Summary 更新用戶資訊
// @Description 更新用戶資訊
// @Tags CMS/User
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "用戶id"
// @Param json_body body validator.CMSUpdateUserBody true "更新欄位"
// @Success 200 {object} model.SuccessResult{data=dto.User} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /cms/user/{user_id} [PATCH]
func (u *CMSUser) UpdateUser(c *gin.Context) {
	var uri validator.UserIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var body validator.CMSUpdateUserBody
	if err := c.ShouldBindJSON(&body); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	user, err := u.userService.UpdateUserByUID(c, *uri.UserID, &dto.UpdateUserParam{
		UserStatus: body.UserStatus,
		Password:   body.Password,
	})
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessResponse(c, user, "success!")
}

func (u *CMSUser) GetUserOrders(c *gin.Context) {

}
