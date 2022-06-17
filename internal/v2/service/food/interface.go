package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Service interface {
	WithTrx(gormTool orm.Tool)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
