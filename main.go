//商品变更通知项目
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/artong0416/goweb/conf"
	"github.com/artong0416/goweb/model"
	"github.com/artong0416/goweb/log"
	"github.com/artong0416/goweb/router"
	"github.com/artong0416/goweb/daemon"
	"github.com/DeanThompson/ginpprof"
)

func main() {
	conf.HookReload = make([]func(), 0)
	// init models
	conf.HookReload = append(conf.HookReload, model.Init)

	//数据库
	model.Init()
	//日志
	//Gin日志定制
	ginLogWriter := &log.GinLoggerWriter{}
	err := ginLogWriter.InitGinLog()

	//初始化路由
	r := gin.New()

	if err == nil {
		r.Use(gin.LoggerWithWriter(ginLogWriter))
	} else {
		log.Log.Error("Gin日志定制失败,不打开Gin日志")
	}
	/*r.Use(gin.LoggerWithWriter(&setting.Logwriter{Mutex: new(sync.Mutex), Prefix: "access", BufferSize: buffer_size}), gin.RecoveryWithWriter(setting.Logwriterfile))*/

	serve := &http.Server{
		Addr:         conf.AppHost,
		Handler:      r,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 10,
	}
	router.Init(r)

	ginpprof.Wrapper(r)


	daemon.SetSignal(serve)
}
