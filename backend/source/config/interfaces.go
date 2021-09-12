package config

type (
	ServerConfigsRepository interface {
		ServerPort() int
		ServerDomain() string
		ServerSecureCookie() bool
	}

	DBConfigsRepository interface {
		DBHost() string
		DBPort() int
		DBUser() string
		DBPassword() string
		DBName() string
		DBConnectionRetryCount() int
		DBConnectionRetryTimeout() int
	}
)
