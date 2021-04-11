package mock

import "time"

const (
	BadLogin       = "BadLogin"
	MissedLogin    = "MissedLogin"
	ExistsLogin    = "ExistsLogin"
	ExistsPassword = "ExistsPassword"

	Format = "2006-01-02 15:04"
)

var (
	Created, _ = time.Parse(Format, "2021-04-11 10:51:27")
)
