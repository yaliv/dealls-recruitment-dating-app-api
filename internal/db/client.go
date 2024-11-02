package db

import (
	"context"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"

	"yaliv/dating-app-api/configs/env"
)

var (
	Client rel.Repository
)

func Open() error {
	adapter, err := postgres.Open(env.DatabaseUrl)
	if err != nil {
		return err
	}

	Client = rel.New(adapter)

	return nil
}

func Close() error {
	if err := Client.Adapter(context.Background()).Close(); err != nil {
		return err
	}

	return nil
}
