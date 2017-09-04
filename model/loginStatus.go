package model

import (
	"time"
)

type LoginStatus struct {
	LoginId   int       `xorm:"not null pk autoincr INT(11)"`
	IpAddress string    `xorm:"not null VARCHAR(20)"`
	TimeIn    time.Time `xorm:"not null TIME"`
	Date      time.Time `xorm:"not null DATE"`
	TimeOut   time.Time `xorm:"not null TIME"`
	StaffId   int       `xorm:"not null INT(11)"`
	Status    int       `xorm:"not null TINYINT(4)"`
}
