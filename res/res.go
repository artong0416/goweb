package res

import "github.com/artong0416/goweb/log"

/*
  #Created by Luheng on 2017/2/14.
  #Arguments: 需要包装的数据
  #Return: 包装后的数据
  #Description: 包装返回数据（正确）
*/
func Return(i interface{}) map[string]interface{} {
	m := make(map[string]interface{}, 3)
	m = map[string]interface{}{"code": "200", "desc": "ok", "data": i}
	return m
}

/*
  #Created by Luheng on 2017/2/14.
  #Arguments: 错误描述string 错误码 int
  #Return: 包装后的数据
  #Description: 包装返回数据（错误）
*/
func ReturnError(msg string, code int) map[string]interface{} {
	m := make(map[string]interface{}, 3)
	m = map[string]interface{}{"code": "500", "desc": "调用失败", "data": nil}
	if len(msg) > 0 {
		m["msg"] = msg
	}

	if code > 0 {
		m["status"] = code
	}

	return m
}

/*
  #Created by Luheng on 2017/8/12.
  #Arguments:
  #Return:
  #Description:通用返回
*/
func ReturnCommon(msg string, data interface{}, status int, desc string) map[string]interface{} {
	m := make(map[string]interface{}, 4)
	m = map[string]interface{}{"msg": msg, "status": status, "data": data}
	log.Log.Info(desc + "返回值:%v", m)
	return m
}