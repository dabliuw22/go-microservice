package controller

import (
	"github.com/gin-gonic/gin"
	"go-microservice/dto"
	"net/http"
)

func HelloWorldRoute(engine *gin.Engine) (*gin.Engine) {
	group := engine.Group("/hello-world")
	{ group.GET("", getHelloWorldHandler) }
	return engine
}

func getHelloWorldHandler(context *gin.Context) {
	context.JSON(http.StatusOK, dto.NewHelloWorldResponse("Hello World"))
}
