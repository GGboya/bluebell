package routes

import (
	"fmt"
	"github.com/GGboya/bluebell/controller"
	"github.com/GGboya/bluebell/logger"
	"github.com/GGboya/bluebell/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	// 初始化gin框架内置校验使用得翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator Trans failed, err:%v\n", err)
		return new(gin.Engine)
	}

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
