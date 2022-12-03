package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	bannerModel "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner_order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner_order/api_update_cms_banner_orders"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	bannerService "github.com/Henry19910227/fitness-go/internal/v2/service/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/service/banner_order"
	"gorm.io/gorm"
)

type resolver struct {
	bannerOrderService banner_order.Service
	bannerService bannerService.Service
}

func New(bannerOrderService banner_order.Service, bannerService bannerService.Service) Resolver {
	return &resolver{bannerOrderService: bannerOrderService, bannerService: bannerService}
}

func (r *resolver) APIUpdateCMSBannerOrders(tx *gorm.DB, input *api_update_cms_banner_orders.Input) (output api_update_cms_banner_orders.Output) {
	defer tx.Rollback()
	// 驗證輸入的 banner id
	bannerIDs := make([]int64, 0)
	for _, bannerOrder := range input.Body.BannerOrders {
		bannerIDs = append(bannerIDs, bannerOrder.BannerID)
	}
	bannerListInput := bannerModel.ListInput{}
	bannerListInput.Wheres = []*whereModel.Where{
		{Query: "banners.id IN (?)", Args: []interface{}{bannerIDs}},
	}
	bannerOutputs, _, err := r.bannerService.Tx(tx).List(&bannerListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(bannerOutputs) != len(input.Body.BannerOrders) {
		output.Set(code.BadRequest, "含有不存在或重複的 banner id")
		return output
	}
	// 刪除原有排序
	if err := r.bannerOrderService.Tx(tx).DeleteAll(); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建新的訓練排序
	tables := make([]*model.Table, 0)
	for _, item := range input.Body.BannerOrders {
		table := model.Table{}
		table.BannerID = util.PointerInt64(item.BannerID)
		table.Seq = util.PointerInt(item.Seq)
		tables = append(tables, &table)
	}
	if err := r.bannerOrderService.Tx(tx).Creates(tables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}