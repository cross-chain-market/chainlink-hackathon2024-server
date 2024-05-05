package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	migrations "github.com/cross-chain-market/chainlink-hackathon2024-server/db"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/config"
)

func New(config *config.Postgres) *bun.DB {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithAddr(config.Address),
		pgdriver.WithUser(config.Username),
		pgdriver.WithPassword(config.Password),
		pgdriver.WithDatabase(config.Database),
		pgdriver.WithApplicationName(config.ServiceName),
		pgdriver.WithInsecure(config.Insecure),
		pgdriver.WithReadTimeout(config.ReadTimeout),
		pgdriver.WithDialTimeout(config.DialTimeout),
	)

	// sql.Register("pg", pgdriver.Driver{})

	db := bun.NewDB(sql.OpenDB(pgconn), pgdialect.New())

	if err := db.Ping(); err != nil {
		slog.Error("failed to run db Ping", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Set max idle connections in pool (default: 2)
	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}

	// Set max time idle pool connections stay open for (default: infinite)
	if config.MaxIdleTime > 0 {
		db.SetConnMaxIdleTime(config.MaxIdleTime)
	}

	// Set max time pool connections can be used for (default: infinite)
	if config.MaxLifeTime > 0 {
		db.SetConnMaxLifetime(config.MaxLifeTime)
	}

	if err := migrateUp(config.DSN()); err != nil {
		slog.Error("failed to run postgres migrations", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return db
}

func migrateUp(uri string) error {
	m, err := newMigrate(uri)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}

func newMigrate(uri string) (*gomigrate.Migrate, error) {
	d, err := iofs.New(migrations.Content, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations: %w", err)
	}

	m, err := gomigrate.NewWithSourceInstance("iofs", d, uri)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrations instance: %w", err)
	}

	return m, nil
}
