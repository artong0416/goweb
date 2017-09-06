/*
作者：陆恒
时间：2017/9/5
功能：
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/go/src/pkg/net/http"
	"github.com/artong0416/goweb/res"
)

// 首页
func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "test/list.html", res.ReturnCommon("test msg", "This is a Test Return!", 200, "TEST"))
}


