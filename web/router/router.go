package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/logic"
)

func StartHttpServer(router *gin.Engine) {
	// favicon.ico
	router.GET("favicon.ico", func(context *gin.Context) {

	})
	routerInit(router)
	_ = router.Run(fmt.Sprintf(":%d", 8888))
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
	}

	// 鉴权
	router.Use(CheckPermission())

	//imS := router.Group("/im")
	//{
	//	imS.GET("/send", sendMessage)
	//}
	//
	userS := router.Group("/user")
	{
		userS.GET("/info/:id", logic.UserInfo)
	}
}
