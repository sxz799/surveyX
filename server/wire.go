//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/utils"
	"gorm.io/gorm"
)

// ServerDeps 包含所有服务器依赖
type ServerDeps struct {
	DB     *gorm.DB
	Config *config.Config
}

func InitializeServer() (*ServerDeps, error) {
	wire.Build(
		config.NewConfig,
		utils.InitDB,
		wire.Struct(new(ServerDeps), "*"),
	)
	return nil, nil
}
