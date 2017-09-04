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
}
