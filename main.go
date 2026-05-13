package main

import (
	"github.com/jfbarkokebas/go-opportunities-api.git/config"
	"github.com/jfbarkokebas/go-opportunities-api.git/router"
)

//após import, comando "go mod tidy" para baixar todas as dependnecias e
// remover as não usadas. Cria go.sum

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
