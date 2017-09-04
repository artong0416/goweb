package model

import (
	"time"
)

type Comment struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	StaffId    int       `xorm:"not null index INT(11)"`
	Text       string    `xorm:"not null TEXT"`
	TimeStamp  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Signature  int       `xorm:"not null TINYINT(2)"`
	DocumentId int       `xorm:"not null INT(11)"`
}
