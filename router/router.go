package router

import (
	"github.com/gin-gonic/gin"
	"go-chat-demo/api"
	"go-chat-demo/service"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//recovery 中间件会恢复（recovery）任何的恐慌（panic）防止程序被迫终止
	r.Use(gin.Recovery(), gin.Logger())
	//Logger日志
	v1 := r.Group("/")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, "SUCCESS")
		})
		//用户注册
		v1.POST("user/register", api.UserRegister)
		//websocket connection
		v1.GET("ws", service.WsHandler)
	}
	return r
}
