package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/ios_version"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
}
