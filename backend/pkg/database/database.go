package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDatabase creates a new database connection based on the driver and dsn
// The dsn is the database connection string, path to the sqlite file or the postgres connection string, not required for memory
func NewDatabase(driver, dsn string) *gorm.DB {
	var dialect gorm.Dialector

	switch driver {
	case "sqlite":
		dialect = sqlite.Open(dsn)
	case "memory":
		dialect = sqlite.Open("file::memory:?cache=shared")
	case "postgres":
		dialect = postgres.Open(dsn)
	default:
		panic("invalid database driver")
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// Database is the database connection, it needs to be set before use
var Database *gorm.DB
