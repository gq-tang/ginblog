package storage

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// DBLogger is a DB wrapper which logs the executed sql queries and their
// duration.
type DBLogger struct {
	*sqlx.DB
}

// Beginx returns a transaction with logging.
func (db *DBLogger) Beginx() (*TxLogger, error) {
	tx, err := db.DB.Beginx()
	return &TxLogger{tx}, err
}

// Query logs the queries executed by the Query method.
func (db *DBLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := db.DB.Query(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// QueryRowx logs the queries executed by the QueryRowx method.
func (db *DBLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	start := time.Now()
	row := db.DB.QueryRowx(query, args...)
	logQuery(query, time.Since(start), args...)
	return row
}

// Exec logs the queries executed by the Exec method.
func (db *DBLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := db.DB.Exec(query, args...)
	logQuery(query, time.Since(start), args...)
	return res, err
}

// Queryx logs the quries executed by the Queryx method.
func (db *DBLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	rows, err := db.DB.Queryx(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// TxLogger logs the executed sql queries and their duration.
type TxLogger struct {
	*sqlx.Tx
}

// Query logs the queries executed by the Query method.
func (tx *TxLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := tx.Tx.Query(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// Queryx logs the queries executed by the Queryx method.
func (tx *TxLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	rows, err := tx.Tx.Queryx(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// QueryRowx logs the queries executed by the QueryRowx method.
func (tx *TxLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	start := time.Now()
	row := tx.Tx.QueryRowx(query, args...)
	logQuery(query, time.Since(start), args...)
	return row
}

// Exec logs the queries executed by the Exec method.
func (tx *TxLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := tx.Tx.Exec(query, args...)
	logQuery(query, time.Since(start), args...)
	return res, err
}

func (tx *TxLogger) Commit() error {
	log.Debug("tx commit")
	return tx.Tx.Commit()
}

func logQuery(query string, duration time.Duration, args ...interface{}) {
	log.WithFields(log.Fields{
		"query":    query,
		"args":     args,
		"duration": duration,
	}).Debug("sql query executed")
}
