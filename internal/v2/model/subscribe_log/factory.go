package subscribe_log

import "fmt"

func GetTypeByIAPType(notificationType string, subtype string) string {
	if notificationType == "SUBSCRIBED" && subtype == "INITIAL_BUY" {
		return InitialBuy
	} else if notificationType == "SUBSCRIBED" && subtype == "RESUBSCRIBE" {
		return Resubscribe
	} else if notificationType == "DID_RENEW" {
		return Renew
	} else if notificationType == "EXPIRED" {
		return Expired
	} else if notificationType == "DID_CHANGE_RENEWAL_PREF" && subtype == "UPGRADE" {
		return Upgrade
	} else if notificationType == "DID_CHANGE_RENEWAL_PREF" && subtype == "DOWNGRADE" {
		return Downgrade
	} else if notificationType == "DID_CHANGE_RENEWAL_PREF" && subtype == "" {
		return DowngradeCancel
	} else if notificationType == "REFUND" {
		return Refund
	} else if notificationType == "DID_CHANGE_RENEWAL_STATUS" && subtype == "AUTO_RENEW_ENABLED" {
		return RenewEnable
	} else if notificationType == "DID_CHANGE_RENEWAL_STATUS" && subtype == "AUTO_RENEW_DISABLED" {
		return RenewDisable
	}
	return Unknown
}

func GetTypeByIABType(notificationType int) string {
	if notificationType == 4 {
		return "initial_buy"
	} else if notificationType == 7 {
		return "resubscribe"
	} else if notificationType == 2 {
		return "renew"
	} else if notificationType == 13 {
		return "expired"
	} else if notificationType == 12 {
		return "refund"
	} else if notificationType == 1 {
		return "renew_enable"
	} else if notificationType == 3 || notificationType == 10 {
		return "renew_disable"
	}
	return "unknown"
}

func GetMsgByIABType(notificationType int) string {
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
