package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	productLabel "github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
	saleItem "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
)

type Table struct {
	IDField
	UserIDField
	SaleTypeField
	SaleIDField
	CourseStatusField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	IntroField
	FoodField
	LevelField
	SuitField
	EquipmentField
	PlaceField
	TrainTargetField
	BodyTargetField
	NoticeField
	PlanCountField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Trainer  *trainer.Table  `json:"trainer,omitempty" gorm:"foreignKey:user_id;references:user_id"` // 教練
	SaleItem *saleItem.Table `json:"sale_item,omitempty" gorm:"foreignKey:id;references:sale_id"`    // 銷售項目
}

func (Table) TableName() string {
	return "courses"
}

// APIGetCMSCoursesOutput /cms/courses [GET] 獲取課表列表 API
type APIGetCMSCoursesOutput struct {
	base.Output
	Data   APIGetCMSCoursesData `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetCMSCoursesData []*struct {
	IDField
	NameField
	CourseStatusField
	ScheduleTypeField
	SaleTypeField
	CreateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.NicknameField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
}

// APIGetCMSCourseOutput /cms/course [GET] 獲取課表詳細 API
type APIGetCMSCourseOutput struct {
	base.Output
	Data *APIGetCMSCourseData `json:"data,omitempty"`
}
type APIGetCMSCourseData struct {
	IDField
	UserIDField
	SaleTypeField
	CourseStatusField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	IntroField
	FoodField
	LevelField
	SuitField
	EquipmentField
	PlaceField
	TrainTargetField
	BodyTargetField
	NoticeField
	CreateAtField
	UpdateAtField
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
}
