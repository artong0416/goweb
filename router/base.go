//Package routers provide the all routers
package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/artong0416/goweb/controller"
)

func Init(r *gin.Engine) {
	//心跳
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": 1, "message": "ok", "data": "pong"})
	})

	r.POST("/test", controller.Test)

/*	r.LoadHTMLGlob("templates/*")
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"title": "Main website",
		})
	})*/

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/test/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test/list.html", gin.H{
			"title": "Main website",
		})
	})

	r.StaticFile("/resource/css/framework.css", "./public/resource/css/framework.css")
	r.StaticFile("/resource/css/main.css", "./public/resource/css/main.css")
	r.StaticFile("/resource/sweetalert/sweetalert.css", "./public/resource/sweetalert/sweetalert.css")

}
