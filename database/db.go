package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	//	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

var Assignment *sqlx.DB

type SSLMode string

const (
	SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

func Connect(host, port, dbname, user, password string, sslMode SSLMode) error {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=%s", host, port, user, password, dbname, SSLModeDisable)
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	Assignment = db
	return migrateStart(db)
}

func migrateStart(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	//NewWithDatabaseInstance returns a new Migrate instance from a source URL
	//and an existing database instance. The source URL scheme is defined by each driver.
	//Use any string that can serve as an identifier during logging as databaseName.
	m, err := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange { //up(): will migrate all the way up
		return err
	}
	return nil
}
