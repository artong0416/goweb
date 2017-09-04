package model

type Setting struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	SettingName string `xorm:"not null VARCHAR(20)"`
	Value       int    `xorm:"not null TINYINT(2)"`
}
