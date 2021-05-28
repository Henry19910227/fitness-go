package handler

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"strconv"
)

var (
	UserTokenPrefix    = "fitness.user.token"
	AdminTokenPrefix   = "fitness.admin.token"

	UserOnlinePrefix   = "fitness.user.online"
)

type sso struct {
	jwtTool tool.JWT
	redisTool tool.Redis
	userSetting setting.User
}

func NewSSO (jwtTool tool.JWT, redisTool tool.Redis, userSetting setting.User) SSO {
	return &sso{jwtTool: jwtTool, redisTool: redisTool, userSetting: userSetting}
}

func (s *sso) GenerateUserToken(uid int64) (string, error) {
	token, err := s.jwtTool.GenerateUserToken(uid)
	if err != nil {
		return "", err
	}
	key := UserTokenPrefix + "." + strconv.Itoa(int(uid))
	if err := s.redisTool.SetEX(key, token, s.jwtTool.GetExpire()); err != nil {
		return "", err
	}
	return token, nil
}


func (s *sso) GenerateAdminToken(uid int64, lv int) (string, error) {
	token, err := s.jwtTool.GenerateAdminToken(uid, lv)
	if err != nil {
		return "", err
	}
	key := AdminTokenPrefix + "." + strconv.Itoa(int(uid))
	if err := s.redisTool.SetEX(key, token, s.jwtTool.GetExpire()); err != nil {
		return "", err
	}
	return token, nil
}

func (s *sso) VerifyUserToken(token string) error {
	return s.authGeneralToken(token, 1)
}

func (s *sso) VerifyLV1AdminToken(token string) error {
	lv, err := s.jwtTool.GetLvByToken(token)
	if err != nil {
		return err
	}
	if lv < 1 {
		return errors.New("error admin")
	}
	return s.authAdminToken(token)
}

func (s *sso) VerifyLV2AdminToken(token string) error {
	lv, err := s.jwtTool.GetLvByToken(token)
	if err != nil {
		return err
	}
	if lv < 2 {
		return errors.New("error admin")
	}
	return s.authAdminToken(token)
}

func (s *sso) ResignAdminToken(token string) error {
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return err
	}
	return s.redisTool.Del(AdminTokenPrefix + "." + strconv.Itoa(int(uid)))
}

func (s *sso) ResignAdminTokenWithUID(uid int64) error {
	return s.redisTool.Del(AdminTokenPrefix + "." + strconv.Itoa(int(uid)))
}

func (s *sso) ResignUserToken(token string) error {
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return err
	}
	return s.redisTool.Del(UserTokenPrefix + "." + strconv.Itoa(int(uid)))
}

func (s *sso) ResignUserTokenWithUID(uid int64) error {
	return s.redisTool.Del(UserTokenPrefix + "." + strconv.Itoa(int(uid)))
}

func (s *sso) authAdminToken(token string) error {
	//從token中取得uid
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return err
	}
	// 該Token未在redis內找到，判定該Token失效，必須重新取得新的Token
	key := AdminTokenPrefix + "." + strconv.Itoa(int(uid))
	currentToken, err := s.redisTool.Get(key)
	if err != nil {
		return err
	}
	// 該Token與當前Token不一致
	if token != currentToken {
		return errors.New("token invalid")
	}
	return nil
}

func (s *sso) authGeneralToken(token string, role int) error {
	//從token中取得uid
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return err
	}
	// 該Token未在redis內找到，判定該Token失效，必須重新取得新的Token
	var key string
	if role == 1 {
		key = UserTokenPrefix + "." + strconv.Itoa(int(uid))
	}
	currentToken, err := s.redisTool.Get(key)
	if err != nil {
		return err
	}
	// 該Token與當前Token不一致
	if token != currentToken {
		return errors.New("token invalid")
	}
	return nil
}