package database

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Migration struct {
	Version  int
	ExecFunc func(pgx.Tx) error
}

// These should be maintained in the order of execution and with
// the correct version number.
var migrations = []Migration{
	{
		Version:  1,
		ExecFunc: InitTables,
	},
}

// This runs SQL queries at the start, if needed.
func RunMigration(conn *pgx.Conn) error {
	// Check if the schema version is the latest
	currentVersion, err := EnsureMigrationsTableExists(conn)
	if err != nil {
		return fmt.Errorf("Failed to get the current schema version: %w", err)
	}

	latestVersion := migrations[len(migrations)-1].Version
	if currentVersion != latestVersion {
		// If the current version is ahead of the latest version, return error.
		if currentVersion > latestVersion {
			return fmt.Errorf("Missing migration file - Current schema version: %d, Latest Version: %d", currentVersion, latestVersion)
		}

		fmt.Printf("Running migrations... (Schema version %d -> %d)\n", currentVersion, latestVersion)

		err = pgx.BeginFunc(context.Background(), conn, func(tx pgx.Tx) error {
			// Run migrations based on the current version.
			for version, migration := range migrations {
				if version > currentVersion {
					err = migration.ExecFunc(tx)
					if err != nil {
						return fmt.Errorf("Failed to run migration version %d: %w", version, err)
					}
				}
			}

			// Update the schema version
			_, err := tx.Exec(context.Background(), "UPDATE migrations SET schema_version = $1", latestVersion)

			return err
		})
		if err != nil {
			return fmt.Errorf("Failed to create migrations table: %w", err)
		}
	}

	return nil
}

// Create a migrations table, if one doesn't already exist and
// returns the current schema version.
func EnsureMigrationsTableExists(conn *pgx.Conn) (int, error) {
	var currentSchemaVersion int

	err := pgx.BeginFunc(context.Background(), conn, func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS migrations (schema_version integer NOT NULL)")
		if err != nil {
			return err
		}

		row := tx.QueryRow(context.Background(), "SELECT schema_version FROM migrations")
		err = row.Scan(&currentSchemaVersion)
		// Add a default entry if no rows found
		if err == pgx.ErrNoRows {
			_, err = tx.Exec(context.Background(), "INSERT INTO migrations VALUES(0)")
			currentSchemaVersion = 0
		}
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		return 0, fmt.Errorf("Failed to create migrations table: %w", err)
	}

	return currentSchemaVersion, nil
}

//go:embed migrations/v1.0_InitTables.sql
var v1_0_InitTables string

func InitTables(tx pgx.Tx) error {
	_, err := tx.Exec(context.Background(), v1_0_InitTables)
	if err != nil {
		return fmt.Errorf("Failed to create initial tables: %w", err)
	}

	return nil
}
