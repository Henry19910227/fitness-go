package user_promote_monthly_statistic

// APIGetCMSUserPromoteStatisticInput /v2/cms/statistic_monthly/user/promote [GET]
type APIGetCMSUserPromoteStatisticInput struct {
	Query APIGetCMSUserPromoteStatisticQuery
}
type APIGetCMSUserPromoteStatisticQuery struct {
	YearRequired
	MonthRequired
}
