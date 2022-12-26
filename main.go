package main

import (
	"Forum-API/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//// 初始化gin实例
	//r := gin.Default()
	//
	//// 注册中间件
	//r.Use(gin.Logger(), gin.Recovery())
	//
	//// 注册路由
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"Hello": "World",
	//	})
	//})
	//
	//// 处理404请求
	//r.NoRoute(func(c *gin.Context) {
	//	// 获取标头信息的Accept信息
	//	acceptString := c.Request.Header.Get("Accept")
	//	if strings.Contains(acceptString, "text/html") {
	//		// 如果是HTML
	//		c.String(http.StatusNotFound, "页面返回 404")
	//	} else {
	//		// 默认返回JSON
	//		c.JSON(http.StatusNotFound, gin.H{
	//			"error_code":    404,
	//			"error_message": "路由未定义，请求url和请求方法是否正确",
	//		})
	//	}
	//})

	// new 一个 gin Engine
	router := gin.New()

	//初始化路由绑定
	bootstrap.SetupRoute(router)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
