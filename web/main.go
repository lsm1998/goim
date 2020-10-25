package main

import (
	"github.com/gin-gonic/gin"
	"web/client"
	"web/router"
)

func main() {
	client.Init()
	router.StartHttpServer(gin.Default())
}
