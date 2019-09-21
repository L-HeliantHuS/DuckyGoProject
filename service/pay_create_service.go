package service

import (
	"DuckyGo/conf"
	"DuckyGo/serializer"
	 "DuckyGo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay"
	"math/rand"
	"time"
)


type CreatePayService struct {
	Price string `form:"price" binding:"required"`

}


func (service *CreatePayService) Create() *serializer.Response {

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

	fmt.Println(url)

	return &serializer.Response{
		Status:    0,
		Data:      gin.H{
			"PayUrl": url,
			"ResultUrl": url.String(),
		},
		Msg:       "",
		Error:     "",
		TimeStamp: 0,
	}
}