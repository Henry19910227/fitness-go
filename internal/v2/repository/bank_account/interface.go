package bank_account

import model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
