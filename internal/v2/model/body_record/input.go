package body_record

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
}

type ListInput struct {
	UserIDOptional
	RecordTypeOptional
	PagingInput
	OrderByInput
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

// APIUpdateBodyRecordInput /body_records [PATCH] 更新體態紀錄
type APIUpdateBodyRecordInput struct {
	Body APIUpdateBodyRecordBody
	Uri APIUpdateBodyRecordUri

}
type APIUpdateBodyRecordBody struct {
	ValueOptional
}
type APIUpdateBodyRecordUri struct {
	IDRequired
}

// APIDeleteBodyRecordInput /body_record/{body_record_id} [DELETE] 刪除體態紀錄
type APIDeleteBodyRecordInput struct {
	Uri APIUpdateBodyRecordUri
}
type APIDeleteBodyRecordUri struct {
	IDRequired
}
