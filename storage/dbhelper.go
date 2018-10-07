package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

// connect database.
func OpenDB(driverName, dsn string) (*DBLogger, error) {
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "dtabase connection error")
	}
	for {
		if err := db.Ping(); err != nil {
			log.Errorf("connet database error: %s,will retry 2 seconds", err)
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	return &DBLogger{db}, nil
}
