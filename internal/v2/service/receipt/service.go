package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/receipt"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository receipt.Repository
}

func New(repository receipt.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error) {
	output, amount, err := s.repository.List(input)
	if err != nil {
		return output, page, err
	}
	page = &paging.Output{}
	page.TotalCount = int(amount)
	page.Page = input.Page
	page.Size = input.Size
	if input.Size != nil {
		page.TotalPage = util.PointerInt(util.Pagination(int(amount), *input.Size))
	}
	return output, page, err
}

func (s *service) Create(item *model.Table) (id int64, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.Create(item)
	return id, err
}

func (s *service) Update(item *model.Table) (err error) {
	input := model.FindInput{}
	input.ID = item.ID
	input.OrderIDField = item.OrderIDField
	input.OriginalTransactionIDField = item.OriginalTransactionIDField
	input.TransactionIDField = item.TransactionIDField
	output, err := s.repository.Find(&input)
	if err != nil {
		return err
	}
	// 將output轉換為table
	var table model.Table
	err = util.Parser(output, &table)
	if err != nil {
		return err
	}
	// 將須更新的值映射到table
	err = util.Parser(item, &table)
	if err != nil {
		return err
	}
	// 更新資料
	err = s.repository.Update(&table)
	return err
}

func (s *service) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	listInput := model.ListInput{}
	listInput.OrderID = item.OrderID
	listInput.OriginalTransactionID = item.OriginalTransactionID
	listInput.TransactionID = item.TransactionID
	outputs, _, err := s.List(&listInput)
	if err != nil {
		return nil, err
	}
	if len(outputs) == 0 {
		id, err := s.Create(item)
		return &id, err
	}
	if len(util.OnNilJustReturnString(item.ReceiptToken, "")) == 0 {
		return id, err
	}
	table := model.Table{}
	table.ID = outputs[0].ID
	table.ReceiptToken = item.ReceiptToken
	err = s.Update(&table)
	return id, err
}
