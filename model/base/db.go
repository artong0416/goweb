//Package Dbshop database
package base

import (
	"sync"
)

/*
  #Created by Luheng on 2017/4/28.
  #Description: 获取数据库连接接口实体
*/
var Db = &ShopModel{BaseModel{Mutex: new(sync.Mutex)}}

type ShopModel struct {
	BaseModel
}
