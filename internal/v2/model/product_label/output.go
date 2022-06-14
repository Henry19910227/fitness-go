package product_label

type Table struct {
	IDField
	NameField
	ProductIDField
	TwdField
}

func (Table) TableName() string {
	return "product_labels"
}
