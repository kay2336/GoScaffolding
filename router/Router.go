package router

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"awesomeProject/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	// 读取配置文件
	conf := utils.GetRouterByViper()
	r := gin.Default()

	// 开始路由
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 设置路由
	v1 := r.Group("v1")
	{
		// 测试v1路由组接口
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "v1/ping OK")
		})

		//登陆注册
		v1.POST("register", controller.Register())
		v1.POST("login", controller.Login())

		auth := v1.Group("/auth")
		auth.Use(middleware.JWT()) //中间件jwt鉴权
		{
			// 测试中间件jwt
			auth.GET("ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, "v1/auth/ping OK")
			})
		}
	}

	// 启动路由
	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		panic(err)
	}
	return r
}
