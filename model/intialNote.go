package model

import (
	"time"
)

type IntialNote struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	StaffId    int       `xorm:"not null INT(11)"`
	Subject    string    `xorm:"not null VARCHAR(60)"`
	Detail     string    `xorm:"not null TEXT"`
	Attachment string    `xorm:"not null VARCHAR(100)"`
	Signature  string    `xorm:"not null VARCHAR(100)"`
	TimeStamp  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Sendto     int       `xorm:"not null INT(11)"`
	Approval   int       `xorm:"not null TINYINT(3)"`
	Archieve   int       `xorm:"not null TINYINT(2)"`
}
