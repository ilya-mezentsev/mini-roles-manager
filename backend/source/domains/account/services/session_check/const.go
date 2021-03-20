package session_check

const (
	missedTokenInHeadersCode        = "missed-token-in-headers"
	missedTokenInHeadersDescription = "No auth token in headers"

	missedTokenInCookiesCode        = "missed-token-in-cookies"
	missedTokenInCookiesDescription = "No auth token in cookies"

	noAccountByTokenCode        = "no-account-by-token"
	noAccountByTokenDescription = "Unable to find account by provided token"
)

const (
	headerTokenKey = "X-RM-Auth-Token"
)
