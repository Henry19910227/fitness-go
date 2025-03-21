package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

// GenerateInput Test Input
type GenerateInput struct {
	DataAmount int
	UserID     []*base.GenerateSetting
	RecordType []*base.GenerateSetting
	Value      []*base.GenerateSetting
}

type FindInput struct {
	IDOptional
}

type ListInput struct {
	UserIDOptional
	RecordTypeOptional
	PagingInput
	OrderByInput
}

type LatestListInput struct {
	UserIDRequired
}

type DeleteInput struct {
	IDOptional
}

// APICreateBodyRecordInput /body_record [POST] 創建體態紀錄 API
type APICreateBodyRecordInput struct {
	UserIDRequired
	Body APICreateBodyRecordBody
}
type APICreateBodyRecordBody struct {
	RecordTypeRequired
	ValueRequired
}

// APIGetBodyRecordsInput /body_records [GET] 獲取體態紀錄 API
type APIGetBodyRecordsInput struct {
	UserIDRequired
	Query APICreateBodyRecordQuery
}
type APICreateBodyRecordQuery struct {
	RecordTypeRequired
	PagingInput
}

// APIGetBodyRecordsLatestInput /body_records/latest [GET] 獲取各類型最新體態紀錄列表
type APIGetBodyRecordsLatestInput struct {
	UserIDRequired
}

// APIUpdateBodyRecordInput /body_records [PATCH] 更新體態紀錄
type APIUpdateBodyRecordInput struct {
	UserIDRequired
	Body APIUpdateBodyRecordBody
	Uri  APIUpdateBodyRecordUri
}
type APIUpdateBodyRecordBody struct {
	ValueOptional
}
type APIUpdateBodyRecordUri struct {
	IDRequired
}

// APIDeleteBodyRecordInput /body_record/{body_record_id} [DELETE] 刪除體態紀錄
type APIDeleteBodyRecordInput struct {
	UserIDRequired
	Uri APIUpdateBodyRecordUri
}
type APIDeleteBodyRecordUri struct {
	IDRequired
}
