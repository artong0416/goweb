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
	"github.com/artong0416/goweb/model"
	"github.com/artong0416/goweb/log"
)

// 首页
func Index(ctx *gin.Context)  {
	r, err := model.GetContryByCode([]string{"ABW","AFG"})
	if err != nil {
		log.Log.Error("查询失败,原因[%s]", err)
		ctx.HTML(http.StatusOK, "test/list.html", res.ReturnCommon("查询失败", "", 200, "TEST"))

	}
	ctx.HTML(http.StatusOK, "test/list.html", res.ReturnCommon("test msg", r, 200, "TEST"))
}


