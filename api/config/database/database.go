package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"

	"github.com/edutav/golang-api-clients/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var database *sql.DB

func Init() {
	var err error
	conf := config.NewEnv()
	connectionString := ""

	value := os.Getenv("AMBIENTE")
	if value == "DOCKER" {
		connectionString = fmt.Sprintf("host=database_postgres port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DBPort, conf.DBUser, conf.DBPass, conf.DBName)
	} else {
		connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPass, conf.DBName)
	}

	database, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic("connectionString error...")
	}

	err = database.Ping()
	if err != nil {
		panic("connectionString error...")
	}

	dirCurrent, _ := os.Getwd()
	dirMigrations := filepath.Join(dirCurrent, "../../migrations")
	log.Print("----------", dirMigrations)

	driver, _ := postgres.WithInstance(database, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://"+dirMigrations,
		"postgres",
		driver,
	)

	m.Steps(2)
}

func CreateConnection() *sql.DB {
	return database
}
