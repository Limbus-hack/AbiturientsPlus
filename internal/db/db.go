package db

import (
	"context"
	"github.com/code7unner/vk-scrapper/config"
	"github.com/jackc/pgx/v4"
)

type DB struct {
	*pgx.Conn
}

func New(ctx context.Context, conf config.CommonEnvConfigs) (*DB, error) {
	parsedConfig, err := pgx.ParseConfig(conf.PostgresDBStr)
	if err != nil {
		return nil, err
	}

	db, err := pgx.ConnectConfig(ctx, parsedConfig)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
