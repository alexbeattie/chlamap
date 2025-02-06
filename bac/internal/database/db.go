package database

import (
    "database/sql"
    "fmt"
    // "os"
    // "path/filepath"
    // "runtime"

    "github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func Initialize(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    return db, nil
}

func checkPostGIS(db *gorm.DB) error {
	var exists bool
	err := db.Raw("SELECT EXISTS(SELECT 1 FROM pg_extension WHERE extname = 'postgis')").Scan(&exists).Error
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("PostGIS extension is not installed in the database")
	}
	return nil
}
func RunMigrations(dsn string, migrationPath string) error {
    // migrationPath := "file:///Users/alexbeattie/Developer/Shafali/bac/internal/database/migration"

    // db, err := sql.Open("postgres", dsn)
    // if err != nil {
    //     return fmt.Errorf("failed to open DB for migrations: %w", err)
    // }
    // defer db.Close()
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return fmt.Errorf("failed to open DB for migrations: %w", err)
    }
    defer db.Close()

    driver, err := migratepg.WithInstance(db, &migratepg.Config{})
    if err != nil {
        return fmt.Errorf("failed to create postgres driver: %w", err)
    }

    m, err := migrate.NewWithDatabaseInstance(migrationPath, "postgres", driver)
    if err != nil {
        return fmt.Errorf("failed to create migrate instance: %w", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to run migrations: %w", err)
    }

    fmt.Println("Migrations applied successfully!")
    return nil
}
