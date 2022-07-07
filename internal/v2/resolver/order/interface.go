package order

import model "github.com/Henry19910227/fitness-go/internal/v2/model/order"

type Resolver interface {
	APIGetCMSOrders(input *model.APIGetCMSOrdersInput) (output model.APIGetCMSOrdersOutput)
}
