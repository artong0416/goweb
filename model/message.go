package model

import (
	"time"
)

type Message struct {
	MessageId int       `xorm:"not null pk autoincr INT(11)"`
	SenderId  int       `xorm:"not null INT(11)"`
	ReciverId int       `xorm:"not null INT(11)"`
	Message   string    `xorm:"not null TEXT"`
	TimeStamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
