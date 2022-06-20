package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type GenerateInput struct {
	DataAmount int
	UserID []*base.GenerateSetting
	RdaID []*base.GenerateSetting
}

type FindInput struct {
	IDField
	preload.Input
}
