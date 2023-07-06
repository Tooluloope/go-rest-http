package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating DB...")
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("error creating migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(

		"file:///migrations",
		"postgres", driver,
	)

	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("error running migration: %w", err)
	}	

	fmt.Println("Migration successful")
	return nil
}