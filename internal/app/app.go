package app

import (
	"context"
	"github.com/code7unner/vk-scrapper/config"
	"github.com/code7unner/vk-scrapper/internal/api/repository"
	"go.uber.org/zap"
)

type App struct {
	Log  *zap.SugaredLogger
	Conf config.CommonEnvConfigs
	Repo *repository.RepoImpl
	Ctx  context.Context
}

func New(log *zap.SugaredLogger, conf config.CommonEnvConfigs, repo *repository.RepoImpl, ctx context.Context) *App {
	return &App{
		Log:  log,
		Conf: conf,
		Repo: repo,
		Ctx:  ctx,
	}
}
