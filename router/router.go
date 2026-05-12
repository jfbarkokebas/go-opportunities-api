package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o Router utilizando as onfigurações default do gin
	router := gin.Default()
	//define uma rota
	initializeRoutes(router)

	router.Run(":8080") //list and serve  on 0.0.0.0:8080
}
