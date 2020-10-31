package router

import (
	"api-gateway/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(router *gin.Engine) {
	favicon(router)
	routerInit(router)
	_ = router.Run(fmt.Sprintf(":%d", 8888))
}

// favicon.ico
func favicon(router *gin.Engine) {
	router.GET("favicon.ico", func(context *gin.Context) {

	})
}

func routerInit(router *gin.Engine) {
	// 来访记录
	// router.Use(AccessRecord())

	// 限流
	router.Use(LimitHandler())

	// 基础服务
	baseS := router.Group("/base")
	{
		baseS.GET("/info", logic.Info)
		baseS.POST("/login", logic.Login)
	}

	// 鉴权
	router.Use(Authorize(), CheckPermission())

	//imS := router.Group("/im")
	//{
	//	imS.GET("/send", sendMessage)
	//}
	//
	userS := router.Group("/user")
	{
		userS.GET("/info", logic.UserInfo)
	}
}
