//Home task (retraining program)
package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/RomaBiliak/lets-go-chat/internal/auth"
	authHttp "github.com/RomaBiliak/lets-go-chat/internal/auth/http"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories/postgre"
	"github.com/RomaBiliak/lets-go-chat/internal/user"
	userHttp "github.com/RomaBiliak/lets-go-chat/internal/user/http"
	httpServer "github.com/RomaBiliak/lets-go-chat/pkg/http"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	connStr := "user="+os.Getenv("PG_USER")+" password="+os.Getenv("PG_PASSWORD")+" dbname="+os.Getenv("PG_DATABASE")+" sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)
	m.Steps(2)
	if err != nil {
		panic(err)
	}

	userRepository := postgre.NewPostgreUserRepository(db)

	userService := user.NewService(userRepository)
	uHttp := userHttp.NewUserHttp(userService)
	http.HandleFunc("/v1/user", uHttp.CreateUser)

	authService := auth.NewService(userRepository)
	aHttp := authHttp.NewAuthHttp(authService)
	http.HandleFunc("/v1/user/login", aHttp.Login)

	httpServer.Start(":8080")
}
