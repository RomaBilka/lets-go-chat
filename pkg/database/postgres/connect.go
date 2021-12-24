package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	Password string
	Database string
	Host     string
}

func Run(config Config) *sql.DB {
	db := connect(config)
	migrateRun(db)
	return db
}

func connect(config Config) *sql.DB {

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}

func migrateRun(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDirPath(),
		"postgres", driver)

	err = m.Up()
	if err == migrate.ErrNoChange {
		fmt.Println("no new migrations")
	}
}

func migrationsDirPath() string {
	_, pathConnect, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot extract files path")
	}
	p := path.Join(path.Dir(pathConnect), "../../../database/migrations")

	if _, err := os.Stat(p); !os.IsNotExist(err) {
		return p
	}
	return "database/migrations"
}
