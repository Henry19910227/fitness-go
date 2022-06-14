package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver course.Resolver
}

func New(resolver course.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSCourses 獲取課表列表
// @Summary 獲取課表列表
// @Description 獲取課表列表
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id query int64 false "課表ID"
// @Param name query string false "課表名稱"
// @Param course_status query int false "課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)"
// @Param sale_type query int false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表)"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} course.APIGetCMSCoursesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/courses [GET]
func (c *controller) GetCMSCourses(ctx *gin.Context) {
	type pagingInput paging.Input
	type orderByInput orderBy.Input
	var query struct {
		model.IDField
		model.NameField
		model.CourseStatusField
		model.SaleTypeField
		pagingInput
		orderByInput
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSCoursesInput{}
	if err := util.Parser(query, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCourses(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSCourse 獲取課表詳細
// @Summary 獲取課表詳細
// @Description 獲取課表詳細
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表ID"
// @Success 200 {object} course.APIGetCMSCourseOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/course/{course_id} [GET]
func (c *controller) GetCMSCourse(ctx *gin.Context) {
	var uri struct {
		model.IDField
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIGetCMSCourseInput{}
	if err := util.Parser(uri, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCourse(&input)
	ctx.JSON(http.StatusOK, output)
}
