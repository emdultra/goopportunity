package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o router utilizando configurações default do gin
	router := gin.Default()

	// Defini rotas
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Inicia o servidor gin na porta 8080
	router.Run(":8080")
}
