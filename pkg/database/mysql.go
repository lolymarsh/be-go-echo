package database

import (
	"fmt"
	"lolymarsh/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDatabaseMySQL(conf *config.DatabaseConfigs) (*sqlx.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	db.SetConnMaxLifetime(time.Duration(conf.ConnectionMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(conf.ConnMaxIdleTime) * time.Second)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	return db, nil
}
