package user_subscribe_info

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
}
