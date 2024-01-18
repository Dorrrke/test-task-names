package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationPath string

	flag.StringVar(&storagePath, "storage-path", "", "addres and port to rin server")
	flag.StringVar(&migrationPath, "migration-path", "", "data base addres")
	flag.Parse()
	if storagePath == "" {

	}
	if migrationPath == "" {

	}

	m, err := migrate.New(
		"file://"+migrationPath,
		storagePath,
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
		}
		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
