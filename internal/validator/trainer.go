package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CreateTrainerForm struct {
	Name             string  `form:"name" binding:"required,max=20" example:"史考特"`                      // 本名 (1~20字元)
	Nickname         string  `form:"nickname" binding:"required,max=20" example:"戰車老師"`                // 暱稱 (1~20字元)
	Skill            []int   `form:"skill" binding:"required,skills" example:"1,3,5"`                     // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	Email            string  `form:"email" binding:"required,email,max=255" example:"jason@gmail.com"`    // 信箱 (最大255字元)
	Phone            string  `form:"phone" binding:"required,startswith=09,len=10" example:"0922244123"`  // 手機
	Address          string  `form:"address" binding:"required,max=200" example:"台北市信義區松智路五段五號"`  // 地址 (最大100字元)
	Intro            string  `form:"intro" binding:"required,max=800" example:"我叫戰車老師"`                // 教練介紹 (1~400字元)
	Experience       int     `form:"experience" binding:"required,max=40" example:"5"`                      // 年資
	Motto            *string  `form:"motto" binding:"omitempty,max=200" example:"戰車老師"`                    // 座右銘 (1~100字元)
	FacebookURL      *string  `form:"facebook_url" binding:"omitempty,max=100" example:"www.facebook.com"`    // 臉書連結
	InstagramURL     *string  `form:"instagram_url" binding:"omitempty,max=100" example:"www.instagram.com"`  // instagram連結
	YoutubeURL       *string  `form:"youtube_url" binding:"omitempty,max=100" example:"www.youtube.com"`      // youtube連結
	CerNames        []string  `form:"certificate_names" binding:"omitempty,max=20" example:"A級教練證照,B級教練證照"` // 證照名稱
	AccountName      string   `form:"account_name" binding:"required,max=40" example:"史考特的帳戶"`              // 帳戶名稱 (1~20字元)
	Account          string   `form:"account" binding:"required,min=6,max=16" example:"090005556789"`           // 帳戶 (6~16字元)
	Branch           string   `form:"branch" binding:"required,max=40" example:"信義分行"`                        // 分行
	BankCode         string   `form:"bank_code" binding:"required,max=40" example:"史考特的帳戶"`                  // 銀行代碼
}

type UpdateTrainerForm struct {
	Nickname         *string  `form:"nickname" binding:"omitempty,max=20" example:"戰車老師"`                // 暱稱 (1~20字元)
	Skill            []int   `form:"skill" binding:"omitempty,skills" example:"1,3,5"`                     // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	Intro            *string  `form:"intro" binding:"omitempty,max=800" example:"我叫戰車老師"`                // 教練介紹 (1~400字元)
	Experience       *int     `form:"experience" binding:"omitempty,max=40" example:"5"`                      // 年資
	Motto            *string  `form:"motto" binding:"omitempty,max=200" example:"戰車老師"`                    // 座右銘 (1~100字元)
	FacebookURL      *string  `form:"facebook_url" binding:"omitempty,max=100" example:"www.facebook.com"`    // 臉書連結
	InstagramURL     *string  `form:"instagram_url" binding:"omitempty,max=100" example:"www.instagram.com"`  // instagram連結
	YoutubeURL       *string  `form:"youtube_url" binding:"omitempty,max=100" example:"www.youtube.com"`      // youtube連結
	DeleteAlbumPhotosIDs []int64 `form:"delete_trainer_album_photos_id" binding:"omitempty" example:"1"` // 證照名稱
	DeleteCerIDs     []int64  `form:"delete_certificate_id"    binding:"omitempty" example:"1"` // 待刪除的證照照片id
	UpdateCerIDs     []int64  `form:"update_certificate_id"    binding:"omitempty" example:"1"` // 待更新的證照照片id
	UpdateCerNames   []string `form:"update_certificate_names" binding:"omitempty,max=40" example:"A級教練證照"` // 待更新的證照名稱
	CreateCerNames   []string `form:"create_certificate_names" binding:"omitempty,max=40" example:"A級教練證照"` // 待新增的證照名稱
}

type TrainerAlbumPhotoIDUri struct {
	PhotoID int64 `uri:"photo_id" binding:"required" example:"1"`
}

type TrainerIDUri struct {
	TrainerID int64 `uri:"user_id" binding:"required" example:"1"`
}

type GetTrainerSummariesQuery struct {
	OrderType *string `form:"order_type" binding:"omitempty,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)-單選
	Page *int `form:"page" binding:"required,min=1" example:"henry"` // 頁數
	Size *int `form:"size" binding:"required,min=1" example:"henry"` // 筆數
}

type GetTrainerCoursesQuery struct {
	OrderType *string `form:"order_type" binding:"omitempty,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)-單選
	Page *int `form:"page" binding:"required,min=1" example:"henry"` // 頁數
	Size *int `form:"size" binding:"required,min=1" example:"henry"` // 筆數
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("skills", Skills)
	}
}

var Skills validator.Func = func(fl validator.FieldLevel) bool {
	return validateSkills(fl, 1,14,6)
}

func validateSkills(fl validator.FieldLevel, min int, max int, maxCount int) bool {
	skills, ok := fl.Field().Interface().([]int)
	if !ok {
		return false
	}
	//檢查是否丟空陣列
	if len(skills) == 0 {
		return false
	}
	//檢查個數是否超過上限
	if len(skills) > maxCount {
		return false
	}

	var maxValue int
	dupMap := make(map[int]int)
	for _, item := range skills {
		//檢查是否重複，沒重複就把新值加入map
		_, ok := dupMap[item]
		if ok {
			return false
		}
		//檢查是否按順序排列
		if item < maxValue {
			return false
		}
		//檢查選項是否在範圍內
		if item < min || item > max {
			return false
		}
		maxValue = item
		dupMap[item] = item
	}
	return true
}
