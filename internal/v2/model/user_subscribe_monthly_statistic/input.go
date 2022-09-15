package user_subscribe_monthly_statistic

// APIGetCMSUserSubscribeStatisticInput /v2/cms/statistic_monthly/user/subscribe [GET]
type APIGetCMSUserSubscribeStatisticInput struct {
	Query APIGetCMSUserSubscribeStatisticQuery
}
type APIGetCMSUserSubscribeStatisticQuery struct {
	YearRequired
	MonthRequired
}
