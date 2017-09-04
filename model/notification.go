package model

type Notification struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	LetterId int    `xorm:"not null INT(11)"`
	Message  string `xorm:"not null TEXT"`
}
