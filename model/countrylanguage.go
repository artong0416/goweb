package model

type Countrylanguage struct {
	Countrycode string  `xorm:"not null pk default '' index CHAR(3)"`
	Language    string  `xorm:"not null pk default '' CHAR(30)"`
	Isofficial  string  `xorm:"not null default 'F' ENUM('T','F')"`
	Percentage  float32 `xorm:"not null default 0.0 FLOAT(4,1)"`
}
