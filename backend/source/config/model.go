package config

type (
	server struct {
		Port         int    `json:"port"`
		Domain       string `json:"domain"`
		SecureCookie bool   `json:"secure_cookie"`
	}

	db struct {
		Host       string `json:"host"`
		Port       int    `json:"port"`
		User       string `json:"user"`
		Password   string `json:"password"`
		DBName     string `json:"db_name"`
		Connection struct {
			RetryCount   int `json:"retry_count"`
			RetryTimeout int `json:"retry_timeout"`
		} `json:"connection"`
	}

	cache struct {
		PermissionLifetime uint `json:"permission_lifetime"`
	}

	settings struct {
		Server server `json:"server"`
		DB     db     `json:"db"`
		Cache  cache  `json:"cache"`
	}
)
