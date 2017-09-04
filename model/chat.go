package model

import (
	"time"
)

type Chat struct {
	Id      int       `xorm:"not null pk autoincr INT(10)"`
	From    string    `xorm:"not null default '' VARCHAR(255)"`
	To      string    `xorm:"not null default '' VARCHAR(255)"`
	Message string    `xorm:"not null TEXT"`
	Sent    time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Recd    int       `xorm:"not null default 0 INT(10)"`
}
