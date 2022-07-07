package receipt

import model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"

type Resolver interface {
	APIGetCMSReceipts(input *model.APIGetCMSReceiptsInput) (output model.APIGetCMSReceiptsOutput)
}
