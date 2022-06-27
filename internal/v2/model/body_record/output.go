package body_record

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

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
	IDField
	UserIDField
	RecordTypeField
	ValueField
	CreateAtField
	UpdateAtField
}
