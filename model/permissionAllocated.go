package model

import (
	"time"
)

type PermissionAllocated struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	RightId   int       `xorm:"not null INT(11)"`
	TimeStamp time.Time `xorm:"not null DATE"`
	StaffId   int       `xorm:"not null INT(11)"`
}
