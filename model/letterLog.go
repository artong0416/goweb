package model

import (
	"time"
)

type LetterLog struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	LetterId  int       `xorm:"not null INT(11)"`
	SenderId  int       `xorm:"not null INT(11)"`
	ReciverId int       `xorm:"not null INT(11)"`
	TimeStamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Received  int       `xorm:"not null TINYINT(2)"`
	Sent      int       `xorm:"not null TINYINT(2)"`
}
