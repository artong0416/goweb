package model

import (
	"github.com/artong0416/goweb/model/base"
	"github.com/pkg/errors"
)

type Country struct {
	Code           string  `json:"Code" xorm:"not null pk default '' CHAR(3)"`
	Name           string  `json:"Name" xorm:"not null default '' CHAR(52)"`
	Continent      string  `json:"Continent" xorm:"not null default 'Asia' ENUM('North America','Africa','Oceania','Antarctica','South America','Asia','Europe')"`
	Region         string  `json:"Region" xorm:"not null default '' CHAR(26)"`
	Surfacearea    float32 `json:"SurfaceArea" xorm:"not null default 0.00 FLOAT(10,2)"`
	Indepyear      int     `json:"IndepYear" xorm:"SMALLINT(6)"`
	Population     int     `json:"Population" xorm:"not null default 0 INT(11)"`
	Lifeexpectancy float32 `json:"LifeExpectancy" xorm:"FLOAT(3,1)"`
	Gnp            float32 `json:"GNP" xorm:"FLOAT(10,2)"`
	Gnpold         float32 `json:"GNPOld" xorm:"FLOAT(10,2)"`
	Localname      string  `json:"LocalName" xorm:"not null default '' CHAR(45)"`
	Governmentform string  `json:"GovernmentForm" xorm:"not null default '' CHAR(45)"`
	Headofstate    string  `json:"HeadOfState" xorm:"CHAR(60)"`
	Capital        int     `json:"Capital" xorm:"INT(11)"`
	Code2          string  `json:"Code2" xorm:"not null default '' CHAR(2)"`
}


func GetContryByCode(codes []string) ([]Country, error) {
	db := base.Db.GetSlave()
	r := make([]Country, 0)

	if len(codes) < 1 {
		return r, nil
	}

	if err := db.Table("country").In("Code", codes).Find(&r); err != nil {
		return nil, errors.Wrap(err, "查询失败")
	}


	return r, nil
}
