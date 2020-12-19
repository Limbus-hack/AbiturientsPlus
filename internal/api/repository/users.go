package repository

import (
	"context"
	"fmt"
	"github.com/code7unner/vk-scrapper/internal/api/model"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type Users interface {
	Create(ctx context.Context, user *model.User) (int, error)
}

type usersImpl struct {
	db  *pgx.Conn
	log *zap.SugaredLogger
}

func NewUsersImpl(db *pgx.Conn, log *zap.SugaredLogger) Users {
	return &usersImpl{db, log}
}

func (u usersImpl) Create(ctx context.Context, user *model.User) (int, error) {
	sql := `insert into users (id, name, last_name, region, prediction, status)` +
		`values ($1, $2, $3, $4, $5, $6)`

	var id int

	err := u.db.QueryRow(
		ctx,
		sql,
		user.ID,
		user.Name,
		user.LastName,
		user.Region,
		user.Prediction,
		user.Status,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u usersImpl) Update(ctx context.Context, id int64, status string) (int, error) {
	sql := `update users set status= $1 where id= $2`

	var UpdatedRows int

	err := u.db.QueryRow(
		ctx,
		sql,
		status,
		id,
	).Scan(&UpdatedRows)
	if err != nil {
		return 0, err
	}

	u.log.Info(fmt.Sprintf("%d rows updated", UpdatedRows))

	return UpdatedRows, nil
}
