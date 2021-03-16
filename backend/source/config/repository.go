package config

type Repository struct {
	s settings
}

func MustNew(settingsFilePath string) Repository {
	return Repository{
		s: mustParse(settingsFilePath),
	}
}

func Default() Repository {
	return Repository{s: settings{
		Server: server{
			Port:   8080,
			Domain: "localhost",
		},
		DB: db{
			Host:     "localhost",
			Port:     5555,
			User:     "roles-manager",
			Password: "password",
			DBName:   "roles_manager",
		},
	}}
}

func (r Repository) ServerPort() int {
	return r.s.Server.Port
}

func (r Repository) ServerDomain() string {
	return r.s.Server.Domain
}

func (r Repository) ServerSecureCookie() bool {
	return r.s.Server.SecureCookie
}

func (r Repository) DBHost() string {
	return r.s.DB.Host
}

func (r Repository) DBPort() int {
	return r.s.DB.Port
}

func (r Repository) DBUser() string {
	return r.s.DB.User
}

func (r Repository) DBPassword() string {
	return r.s.DB.Password
}

func (r Repository) DBName() string {
	return r.s.DB.DBName
}
