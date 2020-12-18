package repository

import (
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type RepoImpl struct {
	Posts Posts
}

func New(db *pgx.Conn, log *zap.SugaredLogger) *RepoImpl {
	return &RepoImpl{
		Posts: NewPostsImpl(db, log),
	}
}
