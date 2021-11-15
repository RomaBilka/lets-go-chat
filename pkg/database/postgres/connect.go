package postgres

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	Password string
	Database string
}

func Run(confid Config) *sql.DB {
	db := connect(confid)
	migrateRun(db)
	return db
}

func connect(config Config) *sql.DB {

	connStr := "user=" + config.User + " password=" + config.Password + " dbname=" + config.Database + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}

func migrateRun(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	err = m.Up()
	if err == migrate.ErrNoChange {
		fmt.Println("no new migrations")
	}
}
