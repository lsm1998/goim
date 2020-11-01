package logic

import (
	"api-gateway/client"
	"context"
	"github.com/gin-gonic/gin"
	"protocols/user"
)

func Login(c *gin.Context) {
	var req user.LoginRequest
	var rsp user.LoginResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	err := client.UserClient.Call(context.Background(), "Login", &req, &rsp)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "服务调用失败",
		})
		return
	}
	c.JSON(200, rsp)
}

func UserInfo(c *gin.Context) {
	var req user.UserInfoRequest
	var rsp user.UserInfoResponse
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	err := client.UserClient.Call(context.Background(), "UserInfo", &req, &rsp)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "服务调用失败",
		})
		return
	}
	c.JSON(200, rsp)
}

func FriendsList(c *gin.Context) {
	var req user.FriendsListRequest
	var rsp user.FriendsListResponse
	uid, _ := c.Get("uid")
	req.UserId = uid.(int64)
	err := client.UserClient.Call(context.Background(), "FriendsList", &req, &rsp)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "服务调用失败",
		})
		return
	}
	c.JSON(200, rsp)
}

func SendMsg(c *gin.Context) {
	client.ImClient.Auth("")
}
