package model

import (
	"time"
)

type NoteLog struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	NoteId    int       `xorm:"not null INT(11)"`
	SenderId  int       `xorm:"not null INT(11)"`
	ReciverId int       `xorm:"not null INT(11)"`
	TimeStamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Received  int       `xorm:"not null TINYINT(4)"`
	Sent      int       `xorm:"not null TINYINT(2)"`
}
