package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_create_trainer_action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_get_trainer_course_actions"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_get_user_action_best_pr"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_update_trainer_action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	actionService "github.com/Henry19910227/fitness-go/internal/v2/service/action"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"gorm.io/gorm"
	"io/ioutil"
	"strconv"
	"strings"
)

type resolver struct {
	actionService   actionService.Service
	courseService   courseService.Service
	coverUploadTool uploader.Tool
	videoUploadTool uploader.Tool
}

func New(actionService actionService.Service, courseService courseService.Service, coverUploadTool uploader.Tool, videoUploadTool uploader.Tool) Resolver {
	return &resolver{actionService: actionService, courseService: courseService, coverUploadTool: coverUploadTool, videoUploadTool: videoUploadTool}
}

func (r *resolver) APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput) {
	actionInput := model.ListInput{}
	actionInput.Source = util.PointerInt(1)
	actionInput.Size = input.Size
	actionInput.Page = input.Page
	actionInput.OrderField = "create_at"
	actionInput.OrderType = order_by.ASC
	datas, page, err := r.actionService.List(&actionInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSActionsData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APICreateCMSAction(input *model.APICreateCMSActionInput) (output model.APICreateCMSActionOutput) {
	table := model.Table{}
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table.Cover = util.PointerString("")
	table.Video = util.PointerString("")
	table.Source = util.PointerInt(1)
	// 儲存動作封面圖
	if input.CoverFile != nil {
		coverNamed, err := r.coverUploadTool.Save(input.CoverFile.Data, input.CoverFile.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Cover = util.PointerString(coverNamed)
	}
	// 儲存動作影片
	if input.VideoFile != nil {
		videoNamed, err := r.videoUploadTool.Save(input.VideoFile.Data, input.VideoFile.Named)
		if err != nil {
			if table.Cover != nil {
				_ = r.coverUploadTool.Delete(*table.Cover)
			}
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Video = util.PointerString(videoNamed)
	}
	result, err := r.actionService.Create(&table)
	if err != nil {
		if table.Cover != nil {
			_ = r.coverUploadTool.Delete(*table.Cover)
		}
		if table.Video != nil {
			_ = r.videoUploadTool.Delete(*table.Video)
		}
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateCMSActionData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateCMSAction(input *model.APIUpdateCMSActionInput) (output base.Output) {
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 儲存動作封面圖
	if input.CoverFile != nil {
		coverNamed, err := r.coverUploadTool.Save(input.CoverFile.Data, input.CoverFile.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Cover = util.PointerString(coverNamed)
	}
	// 儲存動作影片
	if input.VideoFile != nil {
		videoNamed, err := r.videoUploadTool.Save(input.VideoFile.Data, input.VideoFile.Named)
		if err != nil {
			if table.Cover != nil {
				_ = r.coverUploadTool.Delete(*table.Cover)
			}
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Video = util.PointerString(videoNamed)
	}
	// 更新資料
	if err := r.actionService.Update(&table); err != nil {
		if table.Cover != nil {
			_ = r.coverUploadTool.Delete(*table.Cover)
		}
		if table.Video != nil {
			_ = r.videoUploadTool.Delete(*table.Video)
		}
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APICreateUserAction(tx *gorm.DB, input *model.APICreateUserActionInput) (output model.APICreateUserActionOutput) {
	defer tx.Rollback()
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.Source = util.PointerInt(model.SourceUser)
	table.Cover = util.PointerString("")
	table.Video = util.PointerString("")
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	actionOutput, err := r.actionService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if input.Video != nil {
		videoNamed, err := r.videoUploadTool.Save(input.Video.Data, input.Video.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練組
		table := model.Table{}
		table.ID = actionOutput.ID
		table.Video = util.PointerString(videoNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	if input.Cover != nil {
		coverNamed, err := r.coverUploadTool.Save(input.Cover.Data, input.Cover.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練組
		table := model.Table{}
		table.ID = actionOutput.ID
		table.Cover = util.PointerString(coverNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	tx.Commit()
	// parser output
	data := model.APICreateUserActionData{}
	data.ID = actionOutput.ID
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateUserAction(tx *gorm.DB, input *model.APIUpdateUserActionInput) (output model.APIUpdateUserActionOutput) {
	defer tx.Rollback()
	// 查詢動作資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	actionOutput, err := r.actionService.Tx(tx).Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(actionOutput.Source, 0) != model.SourceUser {
		output.Set(code.BadRequest, "非此個人類型動作，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt64(actionOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此動作擁有者，無法修改資源")
		return output
	}
	// 更新動作
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.actionService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if input.Video != nil {
		// 儲存影片
		videoNamed, err := r.videoUploadTool.Save(input.Video.Data, input.Video.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改動作影片
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.Video = util.PointerString(videoNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊影片
		_ = r.videoUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Video, ""))
	}
	if input.Cover != nil {
		// 儲存封面
		coverNamed, err := r.coverUploadTool.Save(input.Cover.Data, input.Cover.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改封面
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.Cover = util.PointerString(coverNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊封面
		_ = r.coverUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Cover, ""))
	}
	tx.Commit()
	// parser output
	// 查詢動作資訊
	findInput = model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	actionOutput, err = r.actionService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIUpdateUserActionData{}
	if err := util.Parser(actionOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetUserActions(input *model.APIGetUserActionsInput) (output model.APIGetUserActionsOutput) {
	// parser input
	sourceList := make([]int, 0)
	if input.Query.Source != nil {
		if strings.Contains(*input.Query.Source, "2") {
			output.Set(code.BadRequest, "搜索內容不可包含教練動作")
			return output
		}
		for _, item := range strings.Split(*input.Query.Source, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			sourceList = append(sourceList, opt)
		}
	} else {
		sourceList = append(sourceList, model.SourceSystem)
		sourceList = append(sourceList, model.SourceUser)
	}

	categoryList := make([]int, 0)
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

	bodyList := make([]int, 0)
	if input.Query.Body != nil {
		for _, item := range strings.Split(*input.Query.Body, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			bodyList = append(bodyList, opt)
		}
	}

	equipmentList := make([]int, 0)
	if input.Query.Equipment != nil {
		for _, item := range strings.Split(*input.Query.Equipment, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			equipmentList = append(equipmentList, opt)
		}
	}
	wheres := make([]*whereModel.Where, 0)
	if len(sourceList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.source IN (?)", Args: []interface{}{sourceList}})
	}
	if len(categoryList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.category IN (?)", Args: []interface{}{categoryList}})
	}
	if len(bodyList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.equipment IN (?)", Args: []interface{}{equipmentList}})
	}
	if len(equipmentList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.body IN (?)", Args: []interface{}{bodyList}})
	}
	// 查詢動作
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Name = input.Query.Name
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	listInput.Wheres = wheres
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "FavoriteAction", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	actionOutputs, page, err := r.actionService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// output 轉換成 data item
	dataItems := make([]*model.APIGetUserActionsItem, 0)
	for _, actionOutput := range actionOutputs {
		dataItem := model.APIGetUserActionsItem{}
		if err := util.Parser(actionOutput, &dataItem); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		dataItem.Favorite = util.PointerInt(0)
		if actionOutput.FavoriteAction != nil {
			dataItem.Favorite = util.PointerInt(1)
		}
		dataItems = append(dataItems, &dataItem)
	}
	// parser output
	data := model.APIGetUserActionsData{}
	if err := util.Parser(dataItems, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetUserActionBestPR(input *api_get_user_action_best_pr.Input) (output api_get_user_action_best_pr.Output) {
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ActionID)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "MaxDistanceRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "MaxRepsRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "MaxRMRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "MaxSpeedRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "MaxWeightRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
		{Field: "MinDurationRecord", Conditions: []interface{}{"user_id = ?", input.UserID}},
	}
	actionOutput, err := r.actionService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_user_action_best_pr.Data{}
	if err := util.Parser(actionOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIDeleteUserAction(input *model.APIDeleteUserActionInput) (output model.APIDeleteUserActionOutput) {
	// 查詢動作資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	actionOutput, err := r.actionService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(actionOutput.Source, 0) != model.SourceUser {
		output.Set(code.BadRequest, "非此個人類型動作，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(actionOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此動作擁有者，無法刪除資源")
		return output
	}
	// 刪除動作
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.ID
	if err := r.actionService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除相關封面
	_ = r.coverUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Cover, ""))
	// 刪除相關影片
	_ = r.videoUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Video, ""))

	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteUserActionVideo(input *model.APIDeleteUserActionVideoInput) (output model.APIDeleteUserActionVideoOutput) {
	// 查詢動作資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	actionOutput, err := r.actionService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(actionOutput.Source, 0) != model.SourceUser {
		output.Set(code.BadRequest, "非此個人類型動作，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(actionOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此動作擁有者，無法刪除資源")
		return output
	}
	// 修改訓練組
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.Video = util.PointerString("")
	if err := r.actionService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除影片
	_ = r.videoUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Video, ""))

	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetUserActionSystemImages() (output model.APIGetUserActionSystemImagesOutput) {
	files, _ := ioutil.ReadDir(util.RootPath() + "/volumes/storage/action/system_image/")
	images := make([]string, 0)
	for _, file := range files {
		images = append(images, file.Name())
	}
	output.Data = &images
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APICreateTrainerAction(tx *gorm.DB, input *api_create_trainer_action.Input) (output api_create_trainer_action.Output) {
	findCourseInput := courseModel.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err := r.courseService.Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	courseStatus := util.OnNilJustReturnInt(courseOutput.CourseStatus, 0)
	if courseStatus != courseModel.Preparing && courseStatus != courseModel.Reject {
		output.Set(code.BadRequest, "該課表不是準備或退審狀態，無法新增動作至該課表")
		return output
	}
	defer tx.Rollback()
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.CourseID = util.PointerInt64(input.Uri.CourseID)
	table.Source = util.PointerInt(model.SourceTrainer)
	table.Cover = util.PointerString("")
	table.Video = util.PointerString("")
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	actionOutput, err := r.actionService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if input.Video != nil {
		videoNamed, err := r.videoUploadTool.Save(input.Video.Data, input.Video.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練組
		table := model.Table{}
		table.ID = actionOutput.ID
		table.Video = util.PointerString(videoNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	if input.Cover != nil {
		coverNamed, err := r.coverUploadTool.Save(input.Cover.Data, input.Cover.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改訓練組
		table := model.Table{}
		table.ID = actionOutput.ID
		table.Cover = util.PointerString(coverNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	tx.Commit()
	// parser output
	data := api_create_trainer_action.Data{}
	data.ID = actionOutput.ID
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetTrainerCourseActions(input *api_get_trainer_course_actions.Input) (output api_get_trainer_course_actions.Output) {
	// parser input
	sourceList := make([]int, 0)
	if input.Query.Source != nil {
		if strings.Contains(*input.Query.Source, "3") {
			output.Set(code.BadRequest, "搜索內容不可包含用戶動作")
			return output
		}
		for _, item := range strings.Split(*input.Query.Source, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			sourceList = append(sourceList, opt)
		}
	} else {
		sourceList = append(sourceList, model.SourceSystem)
		sourceList = append(sourceList, model.SourceTrainer)
	}

	categoryList := make([]int, 0)
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

	bodyList := make([]int, 0)
	if input.Query.Body != nil {
		for _, item := range strings.Split(*input.Query.Body, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			bodyList = append(bodyList, opt)
		}
	}

	equipmentList := make([]int, 0)
	if input.Query.Equipment != nil {
		for _, item := range strings.Split(*input.Query.Equipment, ",") {
			opt, err := strconv.Atoi(item)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			equipmentList = append(equipmentList, opt)
		}
	}
	wheres := make([]*whereModel.Where, 0)
	wheres = append(wheres, &whereModel.Where{Query: "actions.course_id = ? OR actions.course_id IS NULL", Args: []interface{}{input.Uri.CourseID}})
	if len(sourceList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.source IN (?)", Args: []interface{}{sourceList}})
	}
	if len(categoryList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.category IN (?)", Args: []interface{}{categoryList}})
	}
	if len(bodyList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.body IN (?)", Args: []interface{}{bodyList}})
	}
	if len(equipmentList) > 0 {
		wheres = append(wheres, &whereModel.Where{Query: "actions.equipment IN (?)", Args: []interface{}{equipmentList}})
	}
	// 查詢動作
	listInput := model.ListInput{}
	listInput.Name = input.Query.Name
	listInput.Wheres = wheres
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	actionOutputs, page, err := r.actionService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_trainer_course_actions.Data{}
	if err := util.Parser(actionOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateTrainerAction(tx *gorm.DB, input *api_update_trainer_action.Input) (output api_update_trainer_action.Output) {
	defer tx.Rollback()
	// 查詢動作資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{{Field: "Course"}}
	actionOutput, err := r.actionService.Tx(tx).Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(actionOutput.Source, 0) != model.SourceTrainer {
		output.Set(code.BadRequest, "非教練類型動作，無法修改資源")
		return output
	}
	if util.OnNilJustReturnInt64(actionOutput.Course.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此動作擁有者，無法修改資源")
		return output
	}
	// 更新動作
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.actionService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if input.Video != nil {
		// 儲存影片
		videoNamed, err := r.videoUploadTool.Save(input.Video.Data, input.Video.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改動作影片
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.Video = util.PointerString(videoNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊影片
		_ = r.videoUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Video, ""))
	}
	if input.Cover != nil {
		// 儲存封面
		coverNamed, err := r.coverUploadTool.Save(input.Cover.Data, input.Cover.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 修改封面
		table := model.Table{}
		table.ID = util.PointerInt64(input.Uri.ID)
		table.Cover = util.PointerString(coverNamed)
		if err := r.actionService.Tx(tx).Update(&table); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 刪除舊封面
		_ = r.coverUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Cover, ""))
	}
	tx.Commit()
	// 查詢修改後的動作
	findActionInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	actionOutput, err = r.actionService.Find(&findActionInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_update_trainer_action.Data{}
	if err := util.Parser(actionOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIDeleteTrainerAction(input *model.APIDeleteTrainerActionInput) (output model.APIDeleteTrainerActionOutput) {
	// 查詢動作資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	actionOutput, err := r.actionService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(actionOutput.Source, 0) != model.SourceTrainer {
		output.Set(code.BadRequest, "非教練類型動作，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(actionOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此動作擁有者，無法刪除資源")
		return output
	}
	// 刪除動作
	deleteInput := model.DeleteInput{}
	deleteInput.ID = input.Uri.ID
	if err := r.actionService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除相關封面
	_ = r.coverUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Cover, ""))
	// 刪除相關影片
	_ = r.videoUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Video, ""))

	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteTrainerActionVideo(input *model.APIDeleteTrainerActionVideoInput) (output model.APIDeleteTrainerActionVideoOutput) {
	// 查詢動作資訊
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	findInput.Preloads = []*preloadModel.Preload{{Field: "Course"}}
	actionOutput, err := r.actionService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證權限
	if util.OnNilJustReturnInt(actionOutput.Source, 0) != model.SourceTrainer {
		output.Set(code.BadRequest, "非教練類型動作，無法刪除資源")
		return output
	}
	if util.OnNilJustReturnInt64(actionOutput.Course.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此動作擁有者，無法刪除資源")
		return output
	}
	// 修改訓練組
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	table.Video = util.PointerString("")
	if err := r.actionService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除影片
	_ = r.videoUploadTool.Delete(util.OnNilJustReturnString(actionOutput.Video, ""))

	output.Set(code.Success, "success")
	return output
}
