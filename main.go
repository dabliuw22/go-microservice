package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"go-microservice/controller"
)

const PORT_ENV  = "PORT"

func main() {
	engine := gin.Default()
	engine = controller.HelloWorldRoute(engine)
	engine.Run(port())
}

func port() string {
	port := os.Getenv(PORT_ENV)
	if port == "" {
		return ":8080"
	}
	return ":" + port
}
