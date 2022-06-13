package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
)

const Timeout = 5

type DB struct {
	*pg.DB
}

func NewPgDb(pgUrl string) (*DB, error) {
	if pgUrl == "" {
		return nil, errors.New("no PG_URL was provided")
	}

	pgOpts, err := pg.ParseURL(pgUrl)
	if err != nil {
		return nil, err
	}

	pgDB := pg.Connect(pgOpts)

	ctx := context.Background()

	if err := pgDB.Ping(ctx); err != nil {
		//return nil, errors.New("database ping failed")
		return nil, err
	}

	pgDB.WithTimeout(time.Second * time.Duration(Timeout))

	return &DB{pgDB}, nil
}
