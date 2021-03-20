package session

const (
	credentialsNotFoundCode        = "credentials-not-found"
	credentialsNotFoundDescription = "Unable to find account by provided credentials"
)

const (
	cookiePath        = "/"
	cookieMaxAge      = 86400 // 1 day
	cookieHttpOnly    = true
	cookieUnsetValue  = ""
	cookieUnsetMaxAge = 0
)
