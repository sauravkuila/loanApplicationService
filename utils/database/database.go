package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Config is the set of configurable parameters for database
type DBConfig struct {
	DriverName            string
	URL                   string
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
	ConnectionMaxIdleTime time.Duration
}

var db *sql.DB

// InitDatabase is used to initialise the database
func InitDatabase(config DBConfig) error {
	var err error

	// open the database
	db, err = sql.Open(
		config.DriverName,
		config.URL,
	)
	if err != nil {
		return err
	}

	// try to ping
	err = db.Ping()
	if err != nil {
		return err
	}

	// now set the configurations
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxIdleTime(config.ConnectionMaxIdleTime)
	db.SetConnMaxLifetime(config.ConnectionMaxLifetime)

	return nil
}

// Get is used to get the database instance client
func Get() *sql.DB {
	return db
}

// Close is used to close the database instance
func Close() error {
	return db.Close()
}

func GetConnectionsInUse() int {
	return db.Stats().InUse
}
