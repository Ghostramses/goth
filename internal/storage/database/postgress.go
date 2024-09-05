package database

import (
	"context"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

func NewPostgres(ctx context.Context) (*Postgres, error) {
	url := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())

		return nil
	}

	db, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &Postgres{db, &sq, url}, nil
}

func (db *Postgres) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

func (db *Postgres) Close() {
	db.Pool.Close()
}
