package main

import (
	"api-gateway/client"
	"api-gateway/router"
	"github.com/gin-gonic/gin"
)

func main() {
	client.Init()
	router.StartHttpServer(gin.Default())
}
