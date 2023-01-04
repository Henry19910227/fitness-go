package admin

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/admin/api_cms_login"
	"gorm.io/gorm"
)

type Resolver interface {
	APICMSLogin(tx *gorm.DB, input *api_cms_login.Input) (output api_cms_login.Output)
}
