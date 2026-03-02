package common

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Connects to the Postgresql DB
func GetDatabaseConn(context context.Context, config *DatabaseConfig) (*pgx.Conn, error) {
	if config.Username == "" {
		return nil, fmt.Errorf("missing database username")
	}

	if config.Password == "" {
		return nil, fmt.Errorf("missing database password")
	}

	if config.Host == "" {
		return nil, fmt.Errorf("missing database host")
	}

	if config.Port == 0 {
		return nil, fmt.Errorf("missing database port")
	}

	if config.Name == "" {
		return nil, fmt.Errorf("missing database name")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Name)

	conn, err := pgx.Connect(context, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	fmt.Println("Connected to db!")

	return conn, nil
}
