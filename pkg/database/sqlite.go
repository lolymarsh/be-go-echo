package database

import (
	"fmt"
	"lolymarsh/pkg/configs"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabaseSQLite(conf *configs.DatabaseConfigs) (*sqlx.DB, error) {
	dsn := conf.Name // For SQLite, the DSN is typically just the file path or ":memory:" for in-memory DB

	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	// SQLite-specific settings
	db.SetConnMaxLifetime(time.Duration(conf.ConnectionMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(conf.ConnMaxIdleTime) * time.Second)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	return db, nil
}
