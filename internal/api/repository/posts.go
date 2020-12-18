package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const postsCollection = "posts"

type Posts interface {
	GetByID(ctx context.Context, id primitive.ObjectID) ([]Posts, error)
	GetAll(ctx context.Context) ([]Posts, error)
}

type postsImpl struct {
	db  *pgx.Conn
	log *zap.SugaredLogger
}

func NewPostsImpl(db *pgx.Conn, log *zap.SugaredLogger) Posts {
	return &postsImpl{db, log}
}

func (p postsImpl) GetByID(ctx context.Context, id primitive.ObjectID) ([]Posts, error) {
	panic("implement me")
}

func (p postsImpl) GetAll(ctx context.Context) ([]Posts, error) {
	panic("implement me")
}
