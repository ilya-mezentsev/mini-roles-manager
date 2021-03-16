package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/config"
)

func MustGetConnection(c config.DBConfigsRepository) *sqlx.DB {
	db, err := sqlx.Open("postgres", BuildPostgresString(c))
	if err != nil {
		log.Fatalf("Unable to open DB connection: %v", err)
	}

	return db
}

func BuildPostgresString(c config.DBConfigsRepository) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost(),
		c.DBPort(),
		c.DBUser(),
		c.DBPassword(),
		c.DBName(),
	)
}
