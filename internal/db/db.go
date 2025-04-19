package db

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/aventhis/avito-pvz-service/internal/config"
)

// Database представляет собой обертку над sql.DB
type Database struct {
	DB *sql.DB
}

// New создает новое подключение к базе данных
func New(cfg *config.Config) (*Database, error) {
	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("не удалось установить соединение с базой данных: %w", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось проверить соединение с базой данных: %w", err)
	}

	return &Database{DB: db}, nil
}

// Close закрывает соединение с базой данных
func (d *Database) Close() error {
	return d.DB.Close()
}

// RunMigrations запускает миграции базы данных
func (d *Database) RunMigrations() error {
	// Определяем путь к миграциям
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	migrationsPath := fmt.Sprintf("file://%s/migrations", basePath)

	driver, err := postgres.WithInstance(d.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать драйвер миграций: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр миграции: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("не удалось применить миграции: %w", err)
	}

	log.Println("Миграции успешно выполнены")
	return nil
} 