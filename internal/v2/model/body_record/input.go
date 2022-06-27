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
