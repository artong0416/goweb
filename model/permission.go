package model

type Permission struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Value string `xorm:"not null VARCHAR(100)"`
	Name  string `xorm:"not null VARCHAR(100)"`
}
