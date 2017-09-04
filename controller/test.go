package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/artong0416/goweb/log"
	"github.com/artong0416/goweb/res"
)


/*
  #Created by Luheng on 2017/6/1.
  #Description: 测试接口参数
*/
type TestPara struct {
	Body string `form:"body" binding:"required"`
	Sign string `form:"sign"`
}

/*
  #Created by Luheng on 2017/6/1.
  #Arguments:
  #Return:
  #Description: 测试接口Handler
*/
func Test(g *gin.Context) {
	var postpara TestPara
	if err := g.Bind(&postpara); err != nil {
		log.Log.Error("参数错误,参数[%v] 请求方[%s] 原因[%s]", g.Request.PostForm, g.ClientIP(), err.Error())
		g.JSON(http.StatusBadRequest, res.ReturnError("para error! "+err.Error(), -1))
		return
	}
	result := "this is a test"
	g.JSON(http.StatusOK, res.Return(result))

}
