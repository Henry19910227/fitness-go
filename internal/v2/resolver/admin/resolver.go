package admin

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/admin"
	"github.com/Henry19910227/fitness-go/internal/v2/model/admin/api_cms_login"
	"github.com/Henry19910227/fitness-go/internal/v2/service/admin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type resolver struct {
	adminService admin.Service
	redisTool    redis.Tool
	jwtTool      jwt.Tool
}

func New(adminService admin.Service, redisTool redis.Tool, jwtTool jwt.Tool) Resolver {
	return &resolver{adminService: adminService, redisTool: redisTool, jwtTool: jwtTool}
}

func (r *resolver) APICMSLogin(tx *gorm.DB, input *api_cms_login.Input) (output api_cms_login.Output) {
	defer tx.Rollback()
	// 獲取 admin 資訊
	adminListInput := model.ListInput{}
	adminListInput.Email = util.PointerString(input.Body.Email)
	adminListInput.Password = util.PointerString(input.Body.Password)
	adminListInput.Size = util.PointerInt(1)
	adminOutputs, _, err := r.adminService.Tx(tx).List(&adminListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(adminOutputs) == 0 {
		output.Set(code.BadRequest, errors.New("帳號或密碼錯誤").Error())
		return output
	}
	// 修改 admin 資訊
	adminTable := model.Table{}
	adminTable.ID = adminOutputs[0].ID
	adminTable.LastLogin = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	if err := r.adminService.Tx(tx).Update(&adminTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//產生token
	adminID := util.OnNilJustReturnInt64(adminOutputs[0].ID, 0)
	lv := util.OnNilJustReturnInt(adminOutputs[0].Lv, 0)
	token, err := r.jwtTool.GenerateAdminToken(adminID, lv)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//設置token過期時間
	key := jwt.AdminTokenPrefix + "." + strconv.Itoa(int(adminID))
	if err := r.redisTool.SetEX(key, token, r.jwtTool.GetExpire()); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	//Parser Output
	data := api_cms_login.Data{}
	if err := util.Parser(adminOutputs[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &data
	output.Token = util.PointerString(token)
	return output
}
