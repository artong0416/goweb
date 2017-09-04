package model

import (
	"time"
)

type Letter struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	StaffId    int       `xorm:"not null INT(11)"`
	Subject    string    `xorm:"not null VARCHAR(60)"`
	Attachment string    `xorm:"not null VARCHAR(100)"`
	Detail     string    `xorm:"not null TEXT"`
	TimeStamp  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Signature  int       `xorm:"not null TINYINT(2)"`
	Sendto     int       `xorm:"not null INT(11)"`
	Approval   int       `xorm:"not null TINYINT(4)"`
	Archieve   int       `xorm:"not null INT(11)"`
}
