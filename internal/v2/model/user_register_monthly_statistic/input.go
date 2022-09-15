package user_register_monthly_statistic

// APIGetCMSUserRegisterStatisticInput /v2/cms/statistic_monthly/user/register [GET]
type APIGetCMSUserRegisterStatisticInput struct {
	Query APIGetCMSUserRegisterStatisticQuery
}
type APIGetCMSUserRegisterStatisticQuery struct {
	YearRequired
	MonthRequired
}
