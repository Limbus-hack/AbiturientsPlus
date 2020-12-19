package repository

import (
	"github.com/code7unner/vk-scrapper/config"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type RepoImpl struct {
	Posts   Posts
	VkUsers VkUsers
}

func New(db *pgx.Conn, log *zap.SugaredLogger, conf config.CommonEnvConfigs) *RepoImpl {
	return &RepoImpl{
		Posts:   NewPostsImpl(db, log),
		VkUsers: NewVkUserImpl(conf),
	}
}
