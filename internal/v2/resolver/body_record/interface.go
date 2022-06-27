package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
)

type Resolver interface {
	APICreateBodyRecord(input *model.APICreateBodyRecordInput) (output model.APICreateBodyRecordOutput)
	APIGetBodyRecords(input *model.APIGetBodyRecordsInput) (output model.APIGetBodyRecordsOutput)
	APIUpdateBodyRecord(input *model.APIUpdateBodyRecordInput) (output base.Output)
}
