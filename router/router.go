package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o router utilizando configurações default do gin
	router := gin.Default()

	// Defini rotas
	InitializeRoutes(router)

	// Inicia o servidor gin na porta 8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
