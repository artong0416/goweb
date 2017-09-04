package model

type City struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Name        string `xorm:"not null default '' CHAR(35)"`
	Countrycode string `xorm:"not null default '' index CHAR(3)"`
	District    string `xorm:"not null default '' CHAR(20)"`
	Population  int    `xorm:"not null default 0 INT(11)"`
}
