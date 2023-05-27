package router

import (
	"awesomeProject/controller"
	"awesomeProject/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() {
	// 读取配置文件
	conf := utils.GetRouterByViper()

	// 开始路由
	r := gin.Default()

	// 设置路由
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})
		v1.POST("register", controller.Register())
		v1.POST("login", controller.Login())
	}

	// 启动路由
	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		panic(err)
	}
}
