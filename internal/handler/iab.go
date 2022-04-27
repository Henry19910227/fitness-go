package handler

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type iab struct {
	iabTool tool.IAB
	reqTool tool.HttpRequest
}

func NewIAB(iabTool tool.IAB, reqTool tool.HttpRequest) IAB {
	return &iab{iabTool: iabTool, reqTool: reqTool}
}

//V (1) SUBSCRIPTION_RECOVERED - 从帐号保留状态恢复了订阅。
//V (2) SUBSCRIPTION_RENEWED - 续订了处于活动状态的订阅。
//V (3) SUBSCRIPTION_CANCELED - 自愿或非自愿地取消了订阅。如果是自愿取消，在用户取消时发送。
//V (4) SUBSCRIPTION_PURCHASED - 购买了新的订阅。
//(5) SUBSCRIPTION_ON_HOLD - 订阅已进入帐号保留状态（如果已启用）。
//(6) SUBSCRIPTION_IN_GRACE_PERIOD - 订阅已进入宽限期（如果已启用）。
//V (7) SUBSCRIPTION_RESTARTED - 用户已通过 Play > 帐号 > 订阅恢复了订阅。订阅已取消，但在用户恢复时尚未到期。如需了解详情，请参阅 [恢复](/google/play/billing/subscriptions#restore)。
//(8) SUBSCRIPTION_PRICE_CHANGE_CONFIRMED - 用户已成功确认订阅价格变动。
//(9) SUBSCRIPTION_DEFERRED - 订阅的续订时间点已延期。
//V (10) SUBSCRIPTION_PAUSED - 订阅已暂停。
//(11) SUBSCRIPTION_PAUSE_SCHEDULE_CHANGED - 订阅暂停计划已更改。
//V (12) SUBSCRIPTION_REVOKED - 用户在到期时间之前已撤消订阅。
//V (13) SUBSCRIPTION_EXPIRED - 订阅已到期。 V

func (i *iab) ParserIABNotificationType(notificationType int) global.SubscribeLogType {
	if notificationType == 4 {
		return global.InitialBuy
	} else if notificationType == 7 {
		return global.Resubscribe
	} else if notificationType == 2 {
		return global.Renew
	} else if notificationType == 13 {
		return global.Expired
	} else if notificationType == 12 {
		return global.Refund
	} else if notificationType == 1 {
		return global.RenewEnable
	} else if notificationType == 3 || notificationType == 10 {
		return global.RenewDisable
	}
	return global.Unknown
}

func (i *iab) ParserIABNotificationMsg(notificationType int) string {
	dict := make(map[int]string)
	dict[1] = "SUBSCRIPTION_RECOVERED"
	dict[2] = "SUBSCRIPTION_RENEWED"
	dict[3] = "SUBSCRIPTION_CANCELED"
	dict[4] = "SUBSCRIPTION_PURCHASED"
	dict[5] = "SUBSCRIPTION_ON_HOLD"
	dict[6] = "SUBSCRIPTION_IN_GRACE_PERIOD"
	dict[7] = "SUBSCRIPTION_RESTARTED"
	dict[8] = "SUBSCRIPTION_PRICE_CHANGE_CONFIRMED"
	dict[9] = "SUBSCRIPTION_DEFERRED"
	dict[10] = "SUBSCRIPTION_PAUSED"
	dict[11] = "SUBSCRIPTION_PAUSE_SCHEDULE_CHANGED"
	dict[12] = "SUBSCRIPTION_REVOKED"
	dict[13] = "SUBSCRIPTION_EXPIRED"

	result, ok := dict[notificationType]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v %s", notificationType, result)
}

func (i *iab) GetProductsAPI(productID string, purchaseToken string) (*dto.IABProductAPIResponse, error) {
	accessToken, err := i.GetGooglePlayApiAccessToken()
	if err != nil {
		return nil, err
	}
	url := i.iabTool.URL() + "/androidpublisher/v3/applications/" + i.iabTool.PackageName() + "/purchases/products/" + productID + "/tokens/" + purchaseToken
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	result, err := i.reqTool.SendRequest("GET", url, header, nil)
	if err != nil {
		return nil, err
	}
	return dto.NewIABProductAPIResponse(result), nil
}

func (i *iab) GetSubscriptionAPI(productID string, purchaseToken string) (*dto.IABSubscriptionAPIResponse, error) {
	accessToken, err := i.GetGooglePlayApiAccessToken()
	if err != nil {
		return nil, err
	}
	url := i.iabTool.URL() + "/androidpublisher/v3/applications/" + i.iabTool.PackageName() + "/purchases/subscriptions/" + productID + "/tokens/" + purchaseToken
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	result, err := i.reqTool.SendRequest("GET", url, header, nil)
	if err != nil {
		return nil, err
	}
	if errResult, ok := result["error"].(map[string]interface{}); ok {
		msg, ok := errResult["message"].(string)
		if !ok {
			return nil, errors.New("api error")
		}
		return nil, errors.New(msg)
	}
	return dto.NewIABSubscriptionAPIResponse(result), nil
}

func (i *iab) GetGooglePlayApiAccessToken() (string, error) {
	oauthToken, err := i.iabTool.GenerateGoogleOAuth2Token(time.Minute * 30)
	if err != nil {
		return "", err
	}
	param := map[string]interface{}{
		"grant_type": "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"assertion":  oauthToken,
	}
	result, err := i.reqTool.SendRequest("POST", "https://oauth2.googleapis.com/token", nil, param)
	if err != nil {
		return "", err
	}
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("無法取得 access token")
	}
	return accessToken, nil
}