package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/valdirmendesdev/experiences/cmd/experiences/apis"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", apis.GetUser)
	}

	log.Println("Server running...")
	r.Run()

}
