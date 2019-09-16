package serializer

import "DuckyGo/model"

// Lianjia 主要序列化器
type Lianjia struct {
	ID        int    `json:"id"`
	Haddr     string `json:"haddr"`
	Htype     string `json:"htype"`
	Hsize     string `json:"hsize"`
	Horientd  string `json:"horientd"`
	Hstyle    string `json:"hstyle"`
	Hposition string `json:"hposition"`
	Hlike     string `json:"hlike"`
	Htime     string `json:"htime"`
	Hprice    string `json:"hprice"`
	Hmoney    string `json:"hmoney"`
}

// 单个数据序列化器
type LianjiaHouse struct {
	Result Lianjia `json:"result"`
}

// 多个数据序列化器
type LianjiaHouses struct {
	Results []Lianjia `json:"results"`
	Count   int       `json:"count"`
}

// LianjiaResponse 序列化基础模板
func LianjiaResponse(db model.Lianjia) Lianjia {
	return Lianjia{
		ID:        db.Id,
		Haddr:     db.Haddr,
		Htype:     db.Htype,
		Hsize:     db.Hsize,
		Horientd:  db.Horientd,
		Hstyle:    db.Hstyle,
		Hposition: db.Hposition,
		Hlike:     db.Hlike,
		Htime:     db.Htime,
		Hprice:    db.Hprice,
		Hmoney:    db.Hmoney,
	}
}

// LianjiaDataResponse 单个数据序列化响应
func LianjiaDataResponse(db model.Lianjia) LianjiaHouse {
	return LianjiaHouse{
		Result: LianjiaResponse(db),
	}
}

// LianjiaDatasResponse 多个数据序列化响应
func LianjiaDatasResponse(db []model.Lianjia, count int) LianjiaHouses {
	var result []Lianjia
	for _, i := range db {
		result = append(result, LianjiaResponse(i))
	}
	return LianjiaHouses{
		Results: result,
		Count:   count,
	}
}
