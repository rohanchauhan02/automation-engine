package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"  // Import the PostgreSQL driver package
	"github.com/rohanchauhan02/automation-engine/shared/config"

)

// PostgresInterface is an interface represents methods that can be used to get Postgres connection.
type PostgresInterface interface {
	OpenPostgresConn() (*gorm.DB, error)
}

// database is a struct that implements PostgresInterface.
type database struct {
	SharedConfig config.ImmutableConfigInterface
}

func (d *database) OpenPostgresConn() (*gorm.DB, error) {
	log.Println("Opening Postgres connection...")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.SharedConfig.GetDatabaseHost(),
		d.SharedConfig.GetDatabasePort(),
		d.SharedConfig.GetDatabaseUser(),
		d.SharedConfig.GetDatabasePassword(),
		d.SharedConfig.GetDatabaseName(),
	)

	log.Printf("Start open Postgres connection ...%s", connectionString)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Disable table name's pluralization globally
	// db.SingularTable(true)

	// Block global update for all tables to prevent accidentally update all records in the database without a where clause.
	db.BlockGlobalUpdate(true)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(60)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(30 * time.Second)
	// Enable Logger, show detailed log
	db.LogMode(true)

	return db, nil
}

// NewPostgres returns a new instance of PostgresInterface.
func NewPostgres(sharedConfig config.ImmutableConfigInterface) PostgresInterface {
	if sharedConfig == nil {
		panic("[PANIC] sharedConfig is nil, immutable config is required to create a new instance of PostgresInterface")
	}
	return &database{
		SharedConfig: sharedConfig,
	}
}
