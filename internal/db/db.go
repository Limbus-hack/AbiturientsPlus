package db

import (
	"context"
	"fmt"
	"github.com/code7unner/rest-api-template/config"
	"github.com/jackc/pgx/v4"
)

type DB struct {
	*pgx.Conn
}

func New(ctx context.Context, conf config.CommonEnvConfigs) (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.PostgresURL, conf.PostgresPort, conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB)

	db, err := pgx.Connect(ctx, psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close(ctx)

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
