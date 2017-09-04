//Package models provide the basic model of the db
package model

import "github.com/artong0416/goweb/model/base"

func Init() {
	base.Db.InitXorm("db")
	base.Db.GetMaster()
	base.Db.GetSlave()
}
