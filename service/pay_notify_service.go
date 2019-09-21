package service

import (
	"DuckyGo/conf"
	"fmt"
	"net/http"
)

type NotifyPayService struct {

}

func (NotifyPayService) Notify(req *http.Request)  {
	ok, err := conf.AliClient.GetTradeNotification(req)
	fmt.Println(ok, err)
}