package favorite_action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_action"
	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_action"
)

type resolver struct {
	favoriteActionService favorite_action.Service
}

func New(favoriteActionService favorite_action.Service) Resolver {
	return &resolver{favoriteActionService: favoriteActionService}
}

func (r *resolver) APICreateFavoriteAction(input *model.APICreateFavoriteActionInput) (output model.APICreateFavoriteActionOutput) {
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.ActionID = util.PointerInt64(input.Uri.ActionID)
	if err := r.favoriteActionService.Create(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}