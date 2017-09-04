package model

type Country struct {
	Code           string  `xorm:"not null pk default '' CHAR(3)"`
	Name           string  `xorm:"not null default '' CHAR(52)"`
	Continent      string  `xorm:"not null default 'Asia' ENUM('Asia','Europe','North America','Africa','Oceania','Antarctica','South America')"`
	Region         string  `xorm:"not null default '' CHAR(26)"`
	Surfacearea    float32 `xorm:"not null default 0.00 FLOAT(10,2)"`
	Indepyear      int     `xorm:"SMALLINT(6)"`
	Population     int     `xorm:"not null default 0 INT(11)"`
	Lifeexpectancy float32 `xorm:"FLOAT(3,1)"`
	Gnp            float32 `xorm:"FLOAT(10,2)"`
	Gnpold         float32 `xorm:"FLOAT(10,2)"`
	Localname      string  `xorm:"not null default '' CHAR(45)"`
	Governmentform string  `xorm:"not null default '' CHAR(45)"`
	Headofstate    string  `xorm:"CHAR(60)"`
	Capital        int     `xorm:"INT(11)"`
	Code2          string  `xorm:"not null default '' CHAR(2)"`
}
