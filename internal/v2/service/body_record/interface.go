package body_record

import model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"

type Service interface {
	Create(item *model.Table) (output *model.Output, err error)
}
