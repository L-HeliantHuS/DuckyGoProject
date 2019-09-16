package model

type Lianjia struct {
	Id        int `gorm:"primary_key,AUTO_INCREMENT"`
	Haddr     string
	Htype     string
	Hsize     string
	Horientd  string
	Hstyle    string
	Hposition string
	Hlike     string
	Htime     string
	Hprice    string
	Hmoney    string
}
