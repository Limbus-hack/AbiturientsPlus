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
	Update(ctx context.Context, id int, status string) (int, error)
	Retrieve(ctx context.Context, city int, school int) ([]model.User, error)
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

func (u usersImpl) Update(ctx context.Context, id int, status string) (int, error) {
	sql := `update users set status = $1 where id = $2`

	var updatedRows int
	_, err := u.db.Exec(
		ctx,
		sql,
		status,
		id,
	)
	if err != nil {
		return 0, err
	}

	u.log.Info(fmt.Sprintf("%d rows updated", updatedRows))

	return updatedRows, nil
}

func (u usersImpl) Retrieve(ctx context.Context, city int, school int) ([]model.User, error) {
	var queryRows pgx.Rows
	if city != 0 {
		sql := `select * from users where region = $1 and prediction = $2`
		rows, err := u.db.Query(
			ctx,
			sql,
			city,
			school)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		queryRows = rows
	} else {
		sql := `select * from users where prediction = $1`
		rows, err := u.db.Query(
			ctx,
			sql,
			school)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		queryRows = rows
	}

	var users []model.User

	for queryRows.Next() {
		var user model.User
		if err := queryRows.Scan(
			&user.ID,
			&user.Name,
			&user.LastName,
			&user.Region,
			&user.Prediction,
			&user.Status); err != nil {
			return nil, err
		}
		if err := queryRows.Err(); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	u.log.Info("retrieved rows")

	return users, nil
}
