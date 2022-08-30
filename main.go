package main

import (
	"encoding/json"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id int
	Nombre string
	Precio int
	Stock int
	Codigo string
	Publicado bool
	FechaDeCreacion string
}

func main() {

	var p []Producto

	jsonFile, err := os.ReadFile("productos.json")
	if err != nil {
		log.Fatal(err)
	}

	err2 := json.Unmarshal([]byte(jsonFile), &p)
	if err2 != nil {
		log.Fatal(err2)
	}


	// Crea un router con Gin
	router := gin.Default()

	// Captura la solicitud GET de /hello-world
	router.GET("/productos", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Producto": p,
		})
	})
	
	// Corremos nuestro servidor
	router.Run()
}