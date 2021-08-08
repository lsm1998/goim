package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logic/config"
	"logic/route"
)

func main() {
	r := gin.Default()
	route.InitRoute(&r.RouterGroup)
	if err := r.Run(fmt.Sprintf(":%d", config.C.Http.Port)); err != nil {
		panic(err)
	}
}
