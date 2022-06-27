package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "body_images"
}

// APIGetBodyImagesOutput /body_images [GET] 獲取體態照片列表
type APIGetBodyImagesOutput struct {
	base.Output
	Data   APIGetBodyImagesData `json:"data,omitempty"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetBodyImagesData []*struct {
	IDField
	BodyImageField
	WeightField
	CreateAtField
	UpdateAtField
}
