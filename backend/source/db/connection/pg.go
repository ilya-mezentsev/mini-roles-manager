package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/config"
	"time"
)

func MustGetConnection(c config.DBConfigsRepository) *sqlx.DB {
	var (
		db  *sqlx.DB
		err error
	)
	tryNumber := 1
	for {
		db, err = sqlx.Open("postgres", BuildPostgresString(c))
		if err != nil {
			log.Errorf("Unable to open DB connection: %v. try number #%d", err, tryNumber)
			time.Sleep(time.Second * time.Duration(c.DBConnectionRetryTimeout()))
		} else if err = db.Ping(); err != nil {
			log.Errorf("Unable to ping DB: %v. try number #%d", err, tryNumber)
			time.Sleep(time.Second * time.Duration(c.DBConnectionRetryTimeout()))
		} else {
			break
		}

		tryNumber++
		if tryNumber > c.DBConnectionRetryCount() {
			break
		}
	}

	if err != nil {
		log.Fatalf("Unable to create DB connection: %v", err)
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
