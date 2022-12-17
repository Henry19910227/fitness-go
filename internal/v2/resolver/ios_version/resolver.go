package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/ios_version"
	"github.com/Henry19910227/fitness-go/internal/v2/model/ios_version/api_get_ios_version"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/service/ios_version"
)

type resolver struct {
	versionService ios_version.Service
}

func New(versionService ios_version.Service) Resolver {
	return &resolver{versionService: versionService}
}

func (r *resolver) APIGetIOSVersion(input *api_get_ios_version.Input) (output api_get_ios_version.Output) {
	// 查詢最新 ios version
	listInput := model.ListInput{}
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	listInput.Size = util.PointerInt(1)
	versionOutputs, _, err := r.versionService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parse output
	data := api_get_ios_version.Data{}
	if len(versionOutputs) == 0 {
		data.Version = util.PointerString("0")
		output.Set(code.Success, "success")
		output.Data = &data
		return output
	}
	if err := util.Parser(versionOutputs[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
