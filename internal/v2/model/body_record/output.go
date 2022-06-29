package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "body_records"
}

// APICreateBodyRecordOutput /body_record [POST] 創建體態紀錄 API
type APICreateBodyRecordOutput struct {
	base.Output
	Data *APICreateBodyRecordData `json:"data,omitempty"`
}
type APICreateBodyRecordData struct {
	Table
}

// APIGetBodyRecordsOutput /body_records [GET] 獲取動作列表 API
type APIGetBodyRecordsOutput struct {
	base.Output
	Data   APIGetBodyRecordsData `json:"data"`
	Paging *paging.Output        `json:"paging,omitempty"`
}
type APIGetBodyRecordsData []*struct {
	IDField
	RecordTypeField
	ValueField
	CreateAtField
	UpdateAtField
}

// APIGetBodyRecordsLatestOutput /body_records/latest [GET] 獲取各類型最新體態紀錄列表
type APIGetBodyRecordsLatestOutput struct {
	base.Output
	Data APIGetBodyRecordsLatestData `json:"data"`
}
type APIGetBodyRecordsLatestData []*struct {
	IDField
	RecordTypeField
	ValueField
	CreateAtField
	UpdateAtField
}
