package api_migrate_up_to_latest

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

// Output /v2/migrate/up [PUT] 修改並覆蓋餐食 API
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	Version *int  `json:"version,omitempty" example:"1"` // 版本
	Dirty   *bool `json:"dirty,omitempty" example:"0"`   // 是否污染
}
