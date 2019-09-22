package service

import (
	"DuckyGo/conf"
	"DuckyGo/model"
	"DuckyGo/serializer"
	 "DuckyGo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay"
	"log"
	"math/rand"
	"time"
)


type CreatePayService struct {
	Price string `form:"price" binding:"required"`

}


func (service *CreatePayService) Create() *serializer.Response {

	var paydb model.AlipaySuccess

	rand.Seed(time.Now().UnixNano())
	var p = alipay.TradePagePay{}
	p.Subject = "Hello"
	p.OutTradeNo = util.RandStringRunes(20)
	p.TotalAmount = service.Price
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	p.ReturnURL = "http://hls.easy.echosite.cn/alipay/return"
	p.NotifyURL = "http://hls.easy.echosite.cn/alipay/notify"

	url, err := conf.AliClient.TradePagePay(p)

	if err != nil {
		fmt.Println(err)
	}

	paydb.OutTradeNo = p.OutTradeNo

	if err := model.DB.Create(&paydb).Error; err != nil {
		log.Println(err)
		return &serializer.Response{
			Status:    50001,
			Msg:       "数据库操作失败",
		}
	}

	return &serializer.Response{
		Status:    0,
		Data:      gin.H{
			"PayUrl": url,
			"ResultUrl": url.String(),
		},
	}
}