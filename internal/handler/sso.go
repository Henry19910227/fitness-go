package handler

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"strconv"
	"time"
)

var (
	UserTokenPrefix    = "fitness.user.token"
	TrainerTokenPrefix    = "fitness.trainer.token"
	AdminTokenPrefix   = "fitness.admin.token"

	UserOnlinePrefix   = "fitness.user.online"
	TrainerOnlinePrefix   = "fitness.trainer.online"
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

func (s *sso) GenerateTrainerToken(uid int64) (string, error) {
	token, err := s.jwtTool.GenerateTrainerToken(uid)
	if err != nil {
		return "", err
	}
	key := TrainerTokenPrefix + "." + strconv.Itoa(int(uid))
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

func (s *sso) VerifyTrainerToken(token string) error {
	return s.authGeneralToken(token, 2)
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

func (s *sso) RenewOnlineStatus(token string) error {
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return err
	}
	role, err := s.jwtTool.GetRoleByToken(token)
	if err != nil {
		return err
	}
	var key string
	if role == 1 {
		key = UserOnlinePrefix + "." +strconv.Itoa(int(uid))
	}
	if role == 2 {
		key = TrainerOnlinePrefix + "." +strconv.Itoa(int(uid))
	}
	value := time.Now().Format("200601021504")
	return s.redisTool.SetEX(key, value, s.userSetting.GetOnlineExpire())
}

func (s *sso) SetOfflineStatus(token string) error {
	uid, err := s.jwtTool.GetIDByToken(token)
	if err != nil {
		return err
	}
	role, err := s.jwtTool.GetRoleByToken(token)
	if err != nil {
		return err
	}
	if err := s.SetOfflineStatusWithUID(uid, role); err != nil {
		return err
	}
	return nil
}

func (s *sso) SetOfflineStatusWithUID(uid int64, role int) error {
	if role == 1 {
		return s.redisTool.Del(UserOnlinePrefix + "." + strconv.Itoa(int(uid)))
	}
	if role == 2 {
		return s.redisTool.Del(TrainerOnlinePrefix + "." + strconv.Itoa(int(uid)))
	}
	return errors.New("error role type")
}

func (s *sso) GetOnlineDateTime(uid int64) (*time.Time, error) {
	timeStr, err := s.redisTool.Get("fitness.*.online." + strconv.Itoa(int(uid)))
	if err != nil {
		return nil, err
	}
	datetime, err := time.Parse("200601021504", timeStr)
	if err != nil {
		return nil, err
	}
	return &datetime, nil
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
	if role == 2 {
		key = TrainerTokenPrefix + "." + strconv.Itoa(int(uid))
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