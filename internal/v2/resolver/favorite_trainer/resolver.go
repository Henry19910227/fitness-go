package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer/api_create_favorite_trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_trainer"
)

type resolver struct {
	favoriteTrainerService favorite_trainer.Service
}

func New(favoriteCourseService favorite_trainer.Service) Resolver {
	return &resolver{favoriteTrainerService: favoriteCourseService}
}

func (r *resolver) APICreateFavoriteTrainer(input *api_create_favorite_trainer.Input) (output api_create_favorite_trainer.Output) {
	// 查詢教練收藏
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.TrainerID = util.PointerInt64(input.Uri.TrainerID)
	favoriteOutputs, _, err := r.favoriteTrainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(favoriteOutputs) > 0 {
		output.Set(code.BadRequest, "已收藏過該教練")
		return output
	}
	// 新增教練收藏
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.TrainerID = util.PointerInt64(input.Uri.TrainerID)
	if err := r.favoriteTrainerService.Create(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}
