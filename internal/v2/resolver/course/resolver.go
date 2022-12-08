package course

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fcm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_fcm_test"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_get_store_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_get_trainer_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_get_trainer_course_overview"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_update_cms_courses_status"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_update_trainer_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course/api_update_user_course"
	courseStatusLogModel "github.com/Henry19910227/fitness-go/internal/v2/model/course_status_update_log"
	fcmModel "github.com/Henry19910227/fitness-go/internal/v2/model/fcm"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	planModel "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	saleItemModel "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	trainerModel "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	subscribeInfoModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	workoutModel "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_status_update_log"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type resolver struct {
	courseService          courseService.Service
	courseStatusLogService course_status_update_log.Service
	planService            plan.Service
	workoutService         workout.Service
	subscribeInfoService   user_subscribe_info.Service
	saleItemService        sale_item.Service
	trainerService         trainer.Service
	uploadTool             uploader.Tool
	redisTool              redis.Tool
	fcmTool                fcm.Tool
}

func New(courseService courseService.Service, courseStatusLogService course_status_update_log.Service, planService plan.Service,
	workoutService workout.Service, subscribeInfoService user_subscribe_info.Service,
	saleItemService sale_item.Service, trainerService trainer.Service,
	uploadTool uploader.Tool, redisTool redis.Tool, fcmTool fcm.Tool) Resolver {
	return &resolver{courseService: courseService, courseStatusLogService: courseStatusLogService, planService: planService,
		workoutService: workoutService, subscribeInfoService: subscribeInfoService,
		saleItemService: saleItemService, trainerService: trainerService,
		uploadTool: uploadTool, redisTool: redisTool, fcmTool: fcmTool}
}

func (r *resolver) APIFcmTest(input *api_fcm_test.Input) (output api_fcm_test.Output) {
	// 獲取或更新 api token
	apiToken, _ := r.redisTool.Get(r.fcmTool.Key())
	if len(apiToken) == 0 {
		//產出 auth token
		oauthToken, err := r.fcmTool.GenerateGoogleOAuth2Token(time.Hour)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		//獲取API Token
		apiToken, err = r.fcmTool.APIGetGooglePlayToken(oauthToken)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		//儲存API Token
		if err := r.redisTool.SetEX(r.fcmTool.Key(), apiToken, r.fcmTool.GetExpire()); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	// 發出推播
	message := make(map[string]interface{})
	if err := util.Parser(input.Body.Payload, &message); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.fcmTool.APISendMessage(apiToken, message); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetFavoriteCourses(input *model.APIGetFavoriteCoursesInput) (output model.APIGetFavoriteCoursesOutput) {
	// parser input
	param := model.ListInput{}
	param.Joins = []*joinModel.Join{
		{Query: "INNER JOIN favorite_courses ON courses.id = favorite_courses.course_id"},
	}
	param.Wheres = []*whereModel.Where{
		{Query: "favorite_courses.user_id = ?", Args: []interface{}{input.UserID}},
		{Query: "courses.course_status = ?", Args: []interface{}{model.Sale}},
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "ReviewStatistic"},
	}
	param.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("favorite_courses.%s %s", "create_at", order_by.DESC)},
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 執行查詢
	results, page, err := r.courseService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetFavoriteCoursesData{}
	if err := util.Parser(results, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetCMSCourses(ctx *gin.Context, input *model.APIGetCMSCoursesInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Wheres = []*whereModel.Where{
		{Query: "courses.course_status NOT IN (?)", Args: []interface{}{[]int{model.Preparing}}},
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, page, err := r.courseService.List(&param)
	if err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCoursesData{}
	if err := util.Parser(result, &data); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCoursesOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}

func (r *resolver) APIGetCMSCourse(ctx *gin.Context, input *model.APIGetCMSCourseInput) interface{} {
	param := model.FindInput{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, err := r.courseService.Find(&param)
	if err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCourseData{}
	if err := util.Parser(result, &data); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCourseOutput{}
	output.Data = &data
	output.Code = code.Success
	output.Msg = "success!"
	return output
}

func (r *resolver) APIUpdateCMSCoursesStatus(ctx *gin.Context, tx *gorm.DB, input *api_update_cms_courses_status.Input) (output api_update_cms_courses_status.Output) {
	defer tx.Rollback()
	// 查詢課表資訊
	courseListInput := model.ListInput{}
	courseListInput.Wheres = []*whereModel.Where{
		{Query: "courses.id IN (?)", Args: []interface{}{input.Body.IDs}},
	}
	courseListInput.Preloads = []*preloadModel.Preload{
		{Field: "User.Trainer"},
	}
	courseOutputs, _, err := r.courseService.Tx(tx).List(&courseListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseTables := make([]*model.Table, 0)
	logTables := make([]*courseStatusLogModel.Table, 0)
	msgOutputs := make([]*fcmModel.Output, 0)
	for _, courseOutput := range courseOutputs {
		// 準備 course table
		courseTable := model.Table{}
		courseTable.ID = courseOutput.ID
		courseTable.CourseStatus = &input.Body.CourseStatus
		// 準備 log table
		logTable := courseStatusLogModel.Table{}
		logTable.CourseID = courseOutput.ID
		logTable.CourseStatus = util.PointerInt(input.Body.CourseStatus)
		logTable.Comment = util.PointerString("")
		if input.Body.Comment != nil {
			logTable.Comment = input.Body.Comment
		}
		// 準備推播 message
		deviceToken := util.OnNilJustReturnString(courseOutput.UserOnSafe().DeviceToken, "")
		if (input.Body.CourseStatus == model.Sale || input.Body.CourseStatus == model.Reject) && len(deviceToken) > 0 {
			saleTypeMap := map[int]string{
				model.SaleTypeFree:      "免費",
				model.SaleTypeSubscribe: "訂閱",
				model.SaleTypeCharge:    "付費",
			}
			statusDesc := "已經通過審核並上架囉"
			if input.Body.CourseStatus == model.Reject {
				statusDesc = "尚未通過審核"
			}
			trainerName := util.OnNilJustReturnString(courseOutput.UserOnSafe().TrainerOnSafe().Nickname, "")
			courseName := util.OnNilJustReturnString(courseOutput.Name, "")
			saleType := util.OnNilJustReturnInt(courseOutput.SaleType, 0)
			title := "課表審核通知"
			body := fmt.Sprintf("%v教練，你建立的%v課表-%v，%v，請點此打開Fitopia.hub APP確認", trainerName, saleTypeMap[saleType], courseName, statusDesc)
			msgOutput := fcmModel.Output{}
			msgOutput.Message.Token = deviceToken
			msgOutput.Message.Notification.Title = title
			msgOutput.Message.Notification.Body = body
			msgOutput.Message.Data.Title = title
			msgOutput.Message.Data.Body = body
			msgOutputs = append(msgOutputs, &msgOutput)
		}
		courseTables = append(courseTables, &courseTable)
		logTables = append(logTables, &logTable)
	}
	if err := r.courseService.Tx(tx).Updates(courseTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.courseStatusLogService.Tx(tx).Creates(logTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// 獲取或更新 api token
	apiToken, _ := r.redisTool.Get(r.fcmTool.Key())
	if len(apiToken) == 0 {
		//產出 auth token
		oauthToken, _ := r.fcmTool.GenerateGoogleOAuth2Token(time.Hour)
		//獲取API Token
		apiToken, _ = r.fcmTool.APIGetGooglePlayToken(oauthToken)
		//儲存API Token
		_ = r.redisTool.SetEX(r.fcmTool.Key(), apiToken, r.fcmTool.GetExpire())
	}
	// 發出推播
	for _, msgOutput := range msgOutputs {
		message := make(map[string]interface{})
		if err = util.Parser(msgOutput, &message); err != nil {
			logger.Shared().Error(ctx, "APIUpdateCMSCoursesStatus Parser："+err.Error())
			continue
		}
		if err = r.fcmTool.APISendMessage(apiToken, message); err != nil {
			logger.Shared().Error(ctx, "APIUpdateCMSCoursesStatus SendMessage："+err.Error())
			continue
		}
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIUpdateCMSCourseCover(input *model.APIUpdateCMSCourseCoverInput) (output model.APIUpdateCMSCourseCoverOutput) {
	fileNamed, err := r.uploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	table.Cover = util.PointerString(fileNamed)
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}

func (r *resolver) APICreateUserCourse(input *model.APICreateUserCourseInput) (output model.APICreateUserCourseOutput) {
	/** 驗證是否能創建 */
	// 1. 檢查目前創建多計畫課表數量
	courseListInput := model.ListInput{}
	courseListInput.UserID = util.PointerInt64(input.UserID)
	courseListInput.SaleType = util.PointerInt(model.SaleTypePersonal)
	courseListInput.ScheduleType = util.PointerInt(2)
	courseOutputs, _, err := r.courseService.List(&courseListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 2. 檢查目前是否已訂閱
	subscribeListInput := subscribeInfoModel.ListInput{}
	subscribeListInput.UserID = util.PointerInt64(input.UserID)
	subscribeListInput.Page = 1
	subscribeListInput.Size = 1
	subscribeListOutput, _, err := r.subscribeInfoService.List(&subscribeListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 3. 驗證多計畫課表創建是否達上線
	if len(courseOutputs) >= 1 {
		if len(subscribeListOutput) == 0 {
			output.Set(code.PermissionDenied, "多計畫課表創建已達上限(未訂閱)")
			return output
		}
		if util.OnNilJustReturnInt(subscribeListOutput[0].Status, 0) == 0 {
			output.Set(code.PermissionDenied, "多計畫課表創建已達上限(未訂閱)")
			return output
		}
	}
	// 創建課表
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypePersonal)
	table.ScheduleType = util.PointerInt(model.MultiplePlan)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(0)
	table.WorkoutCount = util.PointerInt(0)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateUserCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APICreateUserSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateUserCourseInput) (output model.APICreateUserCourseOutput) {
	defer tx.Rollback()
	//創建單一訓練課表
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypePersonal)
	table.ScheduleType = util.PointerInt(model.SingleWorkout)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(1)
	table.WorkoutCount = util.PointerInt(1)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一計畫
	planTable := planModel.Table{}
	planTable.Name = util.PointerString("單一計畫")
	planTable.CourseID = util.PointerInt64(courseID)
	planTable.WorkoutCount = util.PointerInt(1)
	planID, err := r.planService.Tx(tx).Create(&planTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一訓練
	workoutTable := workoutModel.Table{}
	workoutTable.PlanID = util.PointerInt64(planID)
	workoutTable.Name = util.PointerString("單一訓練")
	workoutTable.Equipment = util.PointerString("")
	workoutTable.StartAudio = util.PointerString("")
	workoutTable.EndAudio = util.PointerString("")
	workoutTable.WorkoutSetCount = util.PointerInt(0)
	_, err = r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	//Parser Output
	data := model.APICreateUserCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetUserPersonalCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput) {
	// 查詢個人課表
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.SaleType = util.PointerInt(model.SaleTypePersonal)
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetUserProgressCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput) {
	// 查詢進行中課表
	listInput := model.ListInput{}
	listInput.Joins = []*joinModel.Join{
		{Query: "LEFT JOIN user_course_statistics ON courses.id = user_course_statistics.course_id"},
	}
	listInput.Wheres = []*whereModel.Where{
		{Query: "user_course_statistics.user_id = ?", Args: []interface{}{input.UserID}},
		{Query: "courses.sale_type IN (?)", Args: []interface{}{[]int{model.SaleTypeFree, model.SaleTypeSubscribe, model.SaleTypeCharge}}},
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "ReviewStatistic"},
		{Field: "SaleItem.ProductLabel"},
	}
	listInput.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("user_course_statistics.%s %s", "update_at", order_by.DESC)},
	}
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetUserChargeCourses(input *model.APIGetUserCoursesInput) (output model.APIGetUserCoursesOutput) {
	// 查詢付費課表
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.SaleType = util.PointerInt(model.SaleTypeCharge)
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	listInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN user_course_assets ON courses.id = user_course_assets.course_id"},
	}
	listInput.Wheres = []*whereModel.Where{
		{Query: "user_course_assets.user_id = ?", Args: []interface{}{input.UserID}},
		{Query: "user_course_assets.available = ?", Args: []interface{}{1}},
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "ReviewStatistic"},
		{Field: "SaleItem.ProductLabel"},
	}
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIDeleteUserCourse(input *model.APIDeleteUserCourseInput) (output model.APIDeleteUserCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != model.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法刪除資源")
		return output
	}
	// 刪除課表
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.ID
	if err := r.courseService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIUpdateUserCourse(input *api_update_user_course.Input) (output api_update_user_course.Output) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != model.SaleTypePersonal {
		output.Set(code.BadRequest, "非個人課表類型，無法修改資源")
		return output
	}
	// 修改課表
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.CourseID)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢課表
	findCourseInput := model.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err = r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_update_user_course.Data{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetUserCourse(input *model.APIGetUserCourseInput) (output model.APIGetUserCourseOutput) {
	// 查詢資料
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "UserCourseStatistic", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "UserCourseAsset", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "FavoriteCourse", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCourseData{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(0)
	data.Favorite = util.PointerInt(0)
	// 獲取是否可訪問狀態
	isAllow, err := r.getAllowAccessStatus(input.UserID, courseOutput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(isAllow)
	// 獲取訂閱狀態
	if courseOutput.FavoriteCourse != nil {
		data.Favorite = util.PointerInt(1)
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetUserCourseStructure(input *model.APIGetUserCourseStructureInput) (output model.APIGetUserCourseStructureOutput) {
	// 檢查課表是否是單一訓練課表
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) != model.SingleWorkout {
		output.Set(code.BadRequest, "只允許查看單一訓練課表")
		return output
	}
	// 查詢課表
	findInput = model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "UserCourseStatistic"},
		{Field: "UserCourseAsset", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "FavoriteCourse", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "Plans"},
		{Field: "Plans.Workouts"},
		{Field: "Plans.Workouts.WorkoutSets"},
		{Field: "Plans.Workouts.WorkoutSets.Action"},
		{Field: "Plans.Workouts.WorkoutSets", Conditions: []interface{}{func(db *gorm.DB) *gorm.DB {
			db = db.Joins("LEFT JOIN workout_set_orders ON workout_sets.id = workout_set_orders.workout_set_id")
			return db.Order("workout_set_orders.seq IS NULL ASC, workout_set_orders.seq ASC, workout_sets.create_at ASC")
		}}},
	}
	courseOutput, err = r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetUserCourseStructureData{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(0)
	data.Favorite = util.PointerInt(0)
	// 獲取是否可訪問狀態
	isAllow, err := r.getAllowAccessStatus(input.UserID, courseOutput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(isAllow)
	// 獲取訂閱狀態
	if courseOutput.FavoriteCourse != nil {
		data.Favorite = util.PointerInt(1)
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetTrainerCourses(input *model.APIGetTrainerCoursesInput) (output model.APIGetTrainerCoursesOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Wheres = []*whereModel.Where{
		{Query: "courses.sale_type IN (?)", Args: []interface{}{[]int{
			model.SaleTypeNone, model.SaleTypeFree,
			model.SaleTypeSubscribe, model.SaleTypeCharge}}},
	}
	listInput.CourseStatus = input.Query.CourseStatus
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetTrainerCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APICreateTrainerCourse(input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput) {
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypeNone)
	table.ScheduleType = util.PointerInt(model.MultiplePlan)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(0)
	table.WorkoutCount = util.PointerInt(0)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateTrainerCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetTrainerCourseOverview(input *api_get_trainer_course_overview.Input) (output api_get_trainer_course_overview.Output) {
	// 查詢資料
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.CourseID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
	}
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證查詢權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非該課表創建者，無法查看課表")
		return output
	}
	// parser output
	data := api_get_trainer_course_overview.Data{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APICreateTrainerSingleWorkoutCourse(tx *gorm.DB, input *model.APICreateTrainerCourseInput) (output model.APICreateTrainerCourseOutput) {
	defer tx.Rollback()
	//創建單一訓練課表
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.SaleType = util.PointerInt(model.SaleTypeNone)
	table.ScheduleType = util.PointerInt(model.SingleWorkout)
	table.CourseStatus = util.PointerInt(model.Preparing)
	table.Category = util.PointerInt(0)
	table.Cover = util.PointerString("")
	table.Intro = util.PointerString("")
	table.Food = util.PointerString("")
	table.Level = util.PointerInt(0)
	table.Suit = util.PointerString("")
	table.Equipment = util.PointerString("")
	table.Place = util.PointerString("")
	table.TrainTarget = util.PointerString("")
	table.BodyTarget = util.PointerString("")
	table.Notice = util.PointerString("")
	table.PlanCount = util.PointerInt(1)
	table.WorkoutCount = util.PointerInt(1)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseID, err := r.courseService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一計畫
	planTable := planModel.Table{}
	planTable.Name = util.PointerString("單一計畫")
	planTable.CourseID = util.PointerInt64(courseID)
	planTable.WorkoutCount = util.PointerInt(1)
	planID, err := r.planService.Tx(tx).Create(&planTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建單一訓練
	workoutTable := workoutModel.Table{}
	workoutTable.PlanID = util.PointerInt64(planID)
	workoutTable.Name = util.PointerString("單一訓練")
	workoutTable.Equipment = util.PointerString("")
	workoutTable.StartAudio = util.PointerString("")
	workoutTable.EndAudio = util.PointerString("")
	workoutTable.WorkoutSetCount = util.PointerInt(0)
	_, err = r.workoutService.Tx(tx).Create(&workoutTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	//Parser Output
	data := model.APICreateTrainerCourseData{}
	data.ID = util.PointerInt64(courseID)
	output.Data = &data
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetTrainerCourse(input *api_get_trainer_course.Input) (output api_get_trainer_course.Output) {
	// 獲取課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.CourseID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
	}
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證
	if input.UserID != util.OnNilJustReturnInt64(courseOutput.UserID, 0) {
		output.Set(code.BadRequest, "非該課表擁有者，無法查看課表資訊")
		return output
	}
	// parser output
	data := api_get_trainer_course.Data{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//查詢退審原因
	if util.OnNilJustReturnInt(data.CourseStatus, 0) == model.Reject {
		logListInput := courseStatusLogModel.ListInput{}
		logListInput.CourseID = util.PointerInt64(input.Uri.CourseID)
		logListInput.CourseStatus = util.PointerInt(model.Reject)
		logListInput.Page = 1
		logListInput.Size = 1
		logListInput.OrderField = "create_at"
		logListInput.OrderType = orderByModel.DESC
		logOutputs, _, err := r.courseStatusLogService.List(&logListInput)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		if len(logOutputs) > 0 {
			data.RejectReason = logOutputs[0].Comment
		}
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateTrainerCourse(tx *gorm.DB, input *api_update_trainer_course.Input) (output api_update_trainer_course.Output) {
	defer tx.Rollback()
	// 查詢關聯課表
	findCourseInput := model.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此課表擁有者，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.CourseStatus, 0) == model.Sale {
		output.Set(code.BadRequest, "此課表銷售中，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt(input.Form.SaleType, 0) == model.SaleTypePersonal {
		output.Set(code.BadRequest, "教練課表無法修改為個人課表")
		return output
	}
	// 驗證輸入參數
	if util.OnNilJustReturnInt(input.Form.SaleType, 0) == model.SaleTypeCharge {
		if input.Form.SaleID == nil {
			output.Set(code.BadRequest, "當銷售類型為付費課表時，SaleID必須帶值")
			return output
		}
		findSaleItemInput := saleItemModel.FindInput{}
		findSaleItemInput.ID = input.Form.SaleID
		saleItemOutput, err := r.saleItemService.Find(&findSaleItemInput)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		if util.OnNilJustReturnInt(saleItemOutput.Type, 0) != model.SaleTypeCharge {
			output.Set(code.BadRequest, "銷售類型不符")
			return output
		}
	}
	// 修改課表
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.CourseID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.courseService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//當課表類型不為付費課表時，sale_id 保持 nil
	findCourseInput = model.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err = r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) != model.SaleTypeCharge {
		if err := r.courseService.Tx(tx).UpdateSaleID(input.Uri.CourseID, nil); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	// 上傳 cover
	if input.Cover != nil {
		// 儲存新 cover
		newCoverNamed, err := r.uploadTool.Save(input.Cover.Data, input.Cover.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改課表
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.CourseID)
		table.Cover = util.PointerString(newCoverNamed)
		if err := r.courseService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊 cover
		_ = r.uploadTool.Delete(util.OnNilJustReturnString(courseOutput.Cover, ""))
	}
	tx.Commit()
	// parser output
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.CourseID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "SaleItem.ProductLabel"},
	}
	courseOutput, err = r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := api_update_trainer_course.Data{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Data = &data
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteTrainerCourse(input *model.APIDeleteTrainerCourseInput) (output model.APIDeleteTrainerCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	// 刪除課表
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.ID
	if err := r.courseService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APISubmitTrainerCourse(input *model.APISubmitTrainerCourseInput) (output model.APISubmitTrainerCourseOutput) {
	// 查詢課表資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt64(courseOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非課表擁有者，無法刪除資源")
		return output
	}
	// 檢查課表狀態
	saleType := util.OnNilJustReturnInt(courseOutput.SaleType, 0)
	if saleType == model.SaleTypeNone {
		output.Set(code.BadRequest, "未設定 sale_type")
		return output
	}
	if saleType == model.SaleTypeCharge && courseOutput.SaleID == nil {
		output.Set(code.BadRequest, "未設定 sale_item_id")
		return output
	}
	courseStatus := util.OnNilJustReturnInt(courseOutput.CourseStatus, 0)
	if courseStatus != model.Preparing && courseStatus != model.Reject {
		output.Set(code.BadRequest, "此課表無法重新送審")
		return output
	}
	// 更新課表
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.CourseStatus = util.PointerInt(model.Reviewing)
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetStoreCourse(input *api_get_store_course.Input) (output api_get_store_course.Output) {
	// 查詢資料
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
		{Field: "UserCourseAsset", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "FavoriteCourse", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_store_course.Data{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(0)
	data.Favorite = util.PointerInt(0)
	// 獲取是否可訪問狀態
	isAllow, err := r.getAllowAccessStatus(input.UserID, courseOutput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(isAllow)
	// 獲取收藏狀態
	if courseOutput.FavoriteCourse != nil {
		data.Favorite = util.PointerInt(1)
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreCourses(input *model.APIGetStoreCoursesInput) (output model.APIGetStoreCoursesOutput) {
	levelList := make([]interface{}, 0)
	if input.Query.Level != nil {
		for _, item := range strings.Split(*input.Query.Level, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			levelList = append(levelList, opt)
		}
	}
	categoryList := make([]interface{}, 0)
	if input.Query.Category != nil {
		for _, item := range strings.Split(*input.Query.Category, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			categoryList = append(categoryList, opt)
		}
	}
	saleTypeList := make([]interface{}, 0)
	if input.Query.SaleType != nil {
		for _, item := range strings.Split(*input.Query.SaleType, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			saleTypeList = append(saleTypeList, opt)
		}
	}
	trainerSexList := make([]interface{}, 0)
	if input.Query.TrainerSex != nil {
		for _, item := range strings.Split(*input.Query.TrainerSex, ",") {
			trainerSexList = append(trainerSexList, item)
		}
	}
	joins := make([]*joinModel.Join, 0)
	wheres := make([]*whereModel.Where, 0)
	orders := make([]*orderByModel.Order, 0)
	if input.Query.Name != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.name LIKE ?", Args: []interface{}{"%" + *input.Query.Name + "%"}})
	}
	if input.Query.Score != nil {
		joins = append(joins, &joinModel.Join{Query: "LEFT JOIN review_statistics AS review ON courses.id = review.course_id"})
		wheres = append(wheres, &whereModel.Where{Query: "FLOOR(review.score_total / review.amount) >= ?", Args: []interface{}{*input.Query.Score}})
	}
	if len(levelList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "courses.level IN (?)", Args: []interface{}{levelList}})
	}
	if len(categoryList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "courses.category IN (?)", Args: []interface{}{categoryList}})
	}
	if input.Query.Suit != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.suit LIKE ?", Args: []interface{}{"%" + *input.Query.Suit + "%"}})
	}
	if input.Query.Equipment != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.equipment LIKE ?", Args: []interface{}{"%" + *input.Query.Equipment + "%"}})
	}
	if input.Query.Place != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.place LIKE ?", Args: []interface{}{"%" + *input.Query.Place + "%"}})
	}
	if input.Query.TrainTarget != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.train_target LIKE ?", Args: []interface{}{"%" + *input.Query.TrainTarget + "%"}})
	}
	if input.Query.BodyTarget != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.body_target LIKE ?", Args: []interface{}{"%" + *input.Query.BodyTarget + "%"}})
	}
	if input.Query.SaleType != nil {
		wheres = append(wheres, &whereModel.Where{Query: "courses.sale_type IN (?)", Args: []interface{}{saleTypeList}})
	}
	if input.Query.TrainerSex != nil {
		joins = append(joins, &joinModel.Join{Query: "INNER JOIN users ON courses.user_id = users.id"})
		wheres = append(wheres, &whereModel.Where{Query: "users.sex IN (?)", Args: []interface{}{trainerSexList}})
	}
	if input.Query.TrainerSkill != nil {
		joins = append(joins, &joinModel.Join{Query: "INNER JOIN trainers ON courses.user_id = trainers.user_id"})
		wheres = append(wheres, &whereModel.Where{Query: "trainers.skill LIKE ?", Args: []interface{}{"%" + *input.Query.TrainerSkill + "%"}})
	}
	if input.Query.OrderField != nil {
		if *input.Query.OrderField == "latest" {
			orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("courses.%s %s", "create_at", order_by.DESC)})
		}
		if *input.Query.OrderField == "popular" {
			joins = append(joins, &joinModel.Join{Query: "LEFT JOIN course_usage_statistics ON courses.id = course_usage_statistics.course_id"})
			orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("course_usage_statistics.%s %s", "total_finish_workout_count", order_by.DESC)})
		}
	} else {
		orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("courses.%s %s", "create_at", order_by.DESC)})
	}
	listInput := model.ListInput{}
	listInput.CourseStatus = util.PointerInt(model.Sale)
	listInput.Wheres = wheres
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
	}
	listInput.Joins = joins
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	listInput.Orders = orders
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetStoreCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreCourseStructure(input *model.APIGetStoreCourseStructureInput) (output model.APIGetStoreCourseStructureOutput) {
	// 檢查課表是否是單一訓練課表
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	courseOutput, err := r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(courseOutput.ScheduleType, 0) != model.SingleWorkout {
		output.Set(code.BadRequest, "只允許查看單一訓練課表")
		return output
	}
	// 查詢課表
	findInput = model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "UserCourseAsset", Conditions: []interface{}{"user_id = ?", input.UserID}}, //判斷是否可訪問權限
		{Field: "FavoriteCourse", Conditions: []interface{}{"user_id = ?", input.UserID}},  //判斷收藏
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
		{Field: "Trainer"},
		{Field: "Plans"},
		{Field: "Plans.Workouts"},
		{Field: "Plans.Workouts.WorkoutSets"},
		{Field: "Plans.Workouts.WorkoutSets.Action"},
		{Field: "Plans.Workouts.WorkoutSets", Conditions: []interface{}{func(db *gorm.DB) *gorm.DB {
			db = db.Joins("LEFT JOIN workout_set_orders ON workout_sets.id = workout_set_orders.workout_set_id")
			return db.Order("workout_set_orders.seq IS NULL ASC, workout_set_orders.seq ASC, workout_sets.create_at ASC")
		}}},
	}
	courseOutput, err = r.courseService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetStoreCourseStructureData{}
	if err := util.Parser(courseOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(0)
	data.Favorite = util.PointerInt(0)
	// 獲取是否可訪問狀態
	isAllow, err := r.getAllowAccessStatus(input.UserID, courseOutput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data.AllowAccess = util.PointerInt(isAllow)
	// 獲取訂閱狀態
	if courseOutput.FavoriteCourse != nil {
		data.Favorite = util.PointerInt(1)
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreTrainerCourses(input *model.APIGetStoreTrainerCoursesInput) (output model.APIGetStoreTrainerCoursesOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.Uri.UserID)
	listInput.CourseStatus = util.PointerInt(model.Sale)
	listInput.SaleType = input.Query.SaleType
	listInput.Wheres = []*whereModel.Where{
		{Query: "courses.sale_type <> ?", Args: []interface{}{model.SaleTypePersonal}}, // not equal
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
	}
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parse output
	data := model.APIGetStoreTrainerCoursesData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreHomePage(input *model.APIGetStoreHomePageInput) (output model.APIGetStoreHomePageOutput) {
	// 查詢最新課表列表
	latestCourseListInput := model.ListInput{}
	latestCourseListInput.CourseStatus = util.PointerInt(model.Sale)
	latestCourseListInput.OrderField = "create_at"
	latestCourseListInput.OrderType = order_by.DESC
	latestCourseListInput.Page = 1
	latestCourseListInput.Size = 5
	latestCourseListInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
	}
	latestCourseOutputs, _, err := r.courseService.List(&latestCourseListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢熱門課表列表
	popularCourseListInput := model.ListInput{}
	popularCourseListInput.CourseStatus = util.PointerInt(model.Sale)
	popularCourseListInput.Page = 1
	popularCourseListInput.Size = 5
	popularCourseListInput.Joins = []*joinModel.Join{
		{Query: "LEFT JOIN course_usage_statistics ON courses.id = course_usage_statistics.course_id"},
	}
	popularCourseListInput.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("course_usage_statistics.%s %s", "total_finish_workout_count", order_by.DESC)},
	}
	popularCourseListInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem.ProductLabel"},
		{Field: "ReviewStatistic"},
	}
	popularCourseOutputs, _, err := r.courseService.List(&popularCourseListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢最新教練
	latestTrainerListInput := trainerModel.ListInput{}
	latestTrainerListInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN users ON users.id = trainers.user_id"},
	}
	latestTrainerListInput.Wheres = []*whereModel.Where{
		{Query: "users.is_deleted = ?", Args: []interface{}{0}},
	}
	latestTrainerListInput.OrderField = "create_at"
	latestTrainerListInput.OrderType = order_by.DESC
	latestTrainerListInput.Page = 1
	latestTrainerListInput.Size = 5
	latestTrainerOutputs, _, err := r.trainerService.List(&latestTrainerListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢熱門教練
	popularTrainerListInput := trainerModel.ListInput{}
	popularTrainerListInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN users ON users.id = trainers.user_id"},
		{Query: "LEFT JOIN trainer_statistics ON trainers.user_id = trainer_statistics.user_id"},
	}
	popularTrainerListInput.Wheres = []*whereModel.Where{
		{Query: "users.is_deleted = ?", Args: []interface{}{0}},
	}
	popularTrainerListInput.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("trainer_statistics.%s %s", "student_count", order_by.DESC)},
	}
	popularTrainerListInput.Page = 1
	popularTrainerListInput.Size = 5
	popularTrainerOutputs, _, err := r.trainerService.List(&popularTrainerListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetStoreHomePageData{}
	if err := util.Parser(latestCourseOutputs, &data.LatestCourses); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(popularCourseOutputs, &data.PopularCourses); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(latestTrainerOutputs, &data.LatestTrainers); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(popularTrainerOutputs, &data.PopularTrainers); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) getAllowAccessStatus(userID int64, courseOutput *model.Output) (isAllow int, err error) {
	isAllow = 0
	// 1.該課表為免費課表
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) == model.SaleTypeFree {
		isAllow = 1
	}
	// 2.該課表為訂閱課表
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) == model.SaleTypeSubscribe {
		infoList := subscribeInfoModel.ListInput{}
		infoList.UserID = util.PointerInt64(userID)
		infoList.Page = 1
		infoList.Size = 1
		infoList.OrderType = order_by.DESC
		infoList.OrderField = "update_at"
		infoOutputs, _, err := r.subscribeInfoService.List(&infoList)
		if err != nil {
			return isAllow, err
		}
		if len(infoOutputs) > 0 {
			if util.OnNilJustReturnInt(infoOutputs[0].Status, 0) == 1 {
				isAllow = 1
			}
		}
	}
	// 3.該課表為付費課表
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) == model.SaleTypeCharge {
		if util.OnNilJustReturnInt(courseOutput.UserCourseAssetOnSafe().Available, 0) == 1 {
			isAllow = 1
		}
	}
	// 4.該課表為個人課表
	if util.OnNilJustReturnInt(courseOutput.SaleType, 0) == model.SaleTypePersonal {
		isAllow = 1
	}
	return isAllow, err
}
