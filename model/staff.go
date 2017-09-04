package model

import (
	"time"
)

type Staff struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	FirstName   string    `xorm:"not null VARCHAR(20)"`
	LastName    string    `xorm:"not null VARCHAR(20)"`
	Cnic        string    `xorm:"not null VARCHAR(20)"`
	Email       string    `xorm:"not null VARCHAR(50)"`
	FatherName  string    `xorm:"not null VARCHAR(20)"`
	Designation string    `xorm:"not null VARCHAR(20)"`
	Signature   string    `xorm:"not null VARCHAR(100)"`
	Password    string    `xorm:"not null VARCHAR(200)"`
	Picture     string    `xorm:"not null VARCHAR(100)"`
	TimeStamp   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Online      int       `xorm:"not null TINYINT(4)"`
}
