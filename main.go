package main

import "github.com/gin-gonic/gin"

func main() {
	// Crea un router con Gin
	router := gin.Default()

	// Captura la solicitud GET de /hello-world
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})
	
	// Corremos nuestro servidor
	router.Run()
}