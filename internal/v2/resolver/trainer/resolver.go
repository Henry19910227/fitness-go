package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	trainerService "github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
)

type resolver struct {
	trainerService trainerService.Service
	uploadTool    uploader.Tool
}

func New(trainerService trainerService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{trainerService: trainerService, uploadTool: uploadTool}
}

func (r *resolver) APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput) {
	fileNamed, err := r.uploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.Avatar = util.PointerString(fileNamed)
	if err := r.trainerService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}
