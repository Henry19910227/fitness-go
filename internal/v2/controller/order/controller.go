package order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver order.Resolver
}

func New(resolver order.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateCourseOrder 創建課表訂單
// @Summary 創建課表訂單
// @Description 創建課表訂單
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body order.APICreateCourseOrderBody true "輸入參數"
// @Success 200 {object} order.APICreateCourseOrderOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/course_order [POST]
func (c *controller) CreateCourseOrder(ctx *gin.Context) {
	input := model.APICreateCourseOrderInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateCourseOrder(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// CreateSubscribeOrder 創建訂閱訂單
// @Summary 創建訂閱訂單
// @Description 創建訂閱訂單
// @Tags 支付_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body order.APICreateSubscribeOrderBody true "輸入參數"
// @Success 200 {object} order.APICreateSubscribeOrderOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/subscribe_order [POST]
func (c *controller) CreateSubscribeOrder(ctx *gin.Context) {
	input := model.APICreateSubscribeOrderInput{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateSubscribeOrder(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSOrders 獲取訂單列表
// @Summary 獲取訂單列表
// @Description 獲取訂單列表
// @Tags CMS訂單管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_id query string false "訂單ID"
// @Param user_id query int64 false "用戶ID"
// @Param type query int false "訂單類型(1:課表購買/2:會員訂閱)"
// @Param order_status query int false "訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} order.APIGetCMSOrdersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/orders [GET]
func (c *controller) GetCMSOrders(ctx *gin.Context) {
	input := model.APIGetCMSOrdersInput{}
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSOrders(&input)
	ctx.JSON(http.StatusOK, output)
}
