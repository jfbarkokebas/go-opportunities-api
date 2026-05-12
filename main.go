package main

import "github.com/jfbarkokebas/go-opportunities-api.git/router"

//após import, comando "go mod tidy" para baixar todas as dependnecias e
// remover as não usadas. Cria go.sum

func main() {
	router.Initialize()
}
