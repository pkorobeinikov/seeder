package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/pkorobeinikov/seeder/seeder"
)

const (
	SeederPgConnStrEnv = "SEEDER_PG_CONNSTR"
)

func Seed(ctx context.Context, cfg seeder.Config) error {
	connStr, found := os.LookupEnv(SeederPgConnStrEnv)
	if !found {
		return errors.New("connection string is not set")
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return errors.Wrap(err, "pgx connect")
	}

	var sql string
	if cfg.SQL != "" {
		sql = cfg.SQL
	} else {
		b, err := os.ReadFile(cfg.File)
		if err != nil {
			return errors.Wrapf(err, "read config file: %s", cfg.File)
		}
		sql = string(b)
	}

	_, err = conn.Exec(ctx, sql)
	if err != nil {
		return errors.Wrap(err, "connection exec query")
	}

	err = conn.Close(ctx)
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "close connection")
	}

	return nil
}

func init() {
	seeder.DefaultRegistry().RegisterSeeder(Seed, "postgres")
}
