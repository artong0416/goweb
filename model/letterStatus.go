package model

import (
	"time"
)

type LetterStatus struct {
	StatusId  int       `xorm:"not null pk autoincr INT(11)"`
	LetterId  int       `xorm:"not null INT(11)"`
	Value     int       `xorm:"not null TINYINT(4)"`
	TimeStamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
