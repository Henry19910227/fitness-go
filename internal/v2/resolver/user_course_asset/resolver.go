package user_course_asset

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	userModel "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset/api_create_cms_course_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset/api_delete_cms_course_user"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_asset"
)

type resolver struct {
	userService            user.Service
	userCourseAssetService user_course_asset.Service
}

func New(userService user.Service, userCourseAssetService user_course_asset.Service) Resolver {
	return &resolver{userService: userService, userCourseAssetService: userCourseAssetService}
}

func (r *resolver) APICreateCMSCourseUsers(input *api_create_cms_course_users.Input) (output api_create_cms_course_users.Output) {
	// 驗證輸入用戶個數
	if len(input.Body.UserIDs) == 0 {
		output.Set(code.BadRequest, "輸入不可為空")
		return output
	}
	if len(input.Body.UserIDs) > 50 {
		output.Set(code.BadRequest, "超過單次新增上限")
		return output
	}
	// 查詢輸入用戶
	userListInput := userModel.ListInput{}
	userListInput.Wheres = []*whereModel.Where{
		{Query: "users.id IN (?)", Args: []interface{}{input.Body.UserIDs}},
	}
	userOutputs, _, err := r.userService.List(&userListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	userMap := make(map[int64]*userModel.Output)
	for _, userOutput := range userOutputs {
		userID := util.OnNilJustReturnInt64(userOutput.ID, 0)
		userMap[userID] = userOutput
	}
	// 驗證輸入用戶有效性
	for _, userID := range input.Body.UserIDs {
		if _, ok := userMap[userID]; !ok {
			output.Set(code.BadRequest, fmt.Sprintf("查無 %v 用戶，請重新確認", userID))
			return output
		}
	}
	// 查詢 asset
	assetListInput := model.ListInput{}
	assetListInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	assetListInput.Available = util.PointerInt(1)
	assetListInput.Wheres = []*whereModel.Where{
		{Query: "user_course_assets.user_id IN (?)", Args: []interface{}{input.Body.UserIDs}},
	}
	assetOutputs, _, err := r.userCourseAssetService.List(&assetListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證用戶贈送資格
	if len(assetOutputs) > 0 {
		userID := util.OnNilJustReturnInt64(assetOutputs[0].UserID, 0)
		output.Set(code.BadRequest, fmt.Sprintf("%v 用戶已有此課表權限", userID))
		return output
	}
	// 新增 asset
	assetTables := make([]*model.Table, 0)
	for _, userID := range input.Body.UserIDs {
		assetTable := model.Table{}
		assetTable.UserID = util.PointerInt64(userID)
		assetTable.CourseID = util.PointerInt64(input.Uri.CourseID)
		assetTable.Available = util.PointerInt(1)
		assetTable.Source = util.PointerInt(model.Gift)
		assetTables = append(assetTables, &assetTable)
	}
	if err := r.userCourseAssetService.Creates(assetTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteCMSCourseUser(input *api_delete_cms_course_user.Input) (output api_delete_cms_course_user.Output) {
	// 查詢 asset
	assetListInput := model.ListInput{}
	assetListInput.UserID = util.PointerInt64(input.Uri.UserID)
	assetListInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	assetListInput.Available = util.PointerInt(1)
	assetOutputs, _, err := r.userCourseAssetService.List(&assetListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(assetOutputs) == 0 {
		output.Set(code.BadRequest, "查無此課表使用者")
		return output
	}
	if util.OnNilJustReturnInt(assetOutputs[0].Source, 0) != model.Gift {
		output.Set(code.BadRequest, "無法刪除非贈送的課表使用者")
		return output
	}
	// 刪除 asset
	deleteAssetInput := model.DeleteInput{}
	deleteAssetInput.ID = util.OnNilJustReturnInt64(assetOutputs[0].ID, 0)
	if err := r.userCourseAssetService.Delete(&deleteAssetInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}
