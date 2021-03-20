package session_check

const (
	missedTokenInCookieCode        = "missed-token-in-cookie"
	missedTokenInCookieDescription = "No auth token in cookies"

	noAccountByTokenCode        = "no-account-by-token"
	noAccountByTokenDescription = "Unable to find account by provided token"
)

const (
	headerTokenKey = "X-RM-Auth-Token"
)
