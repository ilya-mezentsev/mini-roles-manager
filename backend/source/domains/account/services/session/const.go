package session

const (
	credentialsNotFoundCode        = "credentials-not-found"
	credentialsNotFoundDescription = "Unable to find account by provided credentials"

	missedTokenInCookieCode        = "missed-token-in-cookie"
	missedTokenInCookieDescription = "No auth token in cookies"

	noAccountByTokenCode        = "no-account-by-token"
	noAccountByTokenDescription = "Unable to find account by provided token"
)

const (
	cookieTokenKey    = "Roles-Manager-Token"
	cookiePath        = "/"
	cookieMaxAge      = 86400 // 1 day
	cookieHttpOnly    = true
	cookieUnsetValue  = ""
	cookieUnsetMaxAge = 0
)
